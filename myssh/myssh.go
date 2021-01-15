package myssh

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"github.com/romberli/log"
	"golang.org/x/crypto/ssh"
)

//Cli SSH的连接参数结构体
type Cli struct {
	//IP地址
	IP string
	//用户名
	User string
	//密码
	Password string
	//端口号
	Port int
	//ssh客户端，连接完毕后会生成一个指针，默认为空
	Client *ssh.Client
	//sftp客户端，连接完毕后会生成一个指针，默认为空
	SftpClient *sftp.Client
}

//NewCli 创建一个新的Cli结构体，指定连接参数，client和LastResult为空
//@param ip IP地址
//@param User 用户名
//@param password 密码
//@param port 端口号,默认22
func NewCli(ip string, User string, password string, port ...int) *Cli {
	cli := new(Cli)
	cli.IP = ip
	cli.User = User
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}

	return cli
}

//connect 创建连接得到c.Client
func (c *Cli) connect() error {
	config := ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	fmt.Println("通讯地址；", addr)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	c.Client = sshClient
	// get auth method
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(c.Password))
	// create sftp Client
	if _, err = sftp.NewClient(c.Client); err != nil {
		return err
	}
	c.SftpClient, _ = sftp.NewClient(c.Client)

	return nil
}

//StartConnect 创建测试
func (c *Cli) StartConnect() (err error) {
	if c.Client == nil {
		if err := c.connect(); err != nil {
			log.Warnf("建立连接失败，error:%v", err)
			return err
		}
	}

	log.Info("成功建立连接")
	return nil
}

//Run 用于执行命令
func (c *Cli) Run(shell string) (result string, err error) {
	if c.Client != nil {
		s, err := c.Client.NewSession()
		defer s.Close()
		if err != nil {
			log.Warnf("建立对话失败%v", err)
			return "", err
		}

		if err == nil {
			buf, err := s.CombinedOutput(shell)
			result = string(buf)
			fmt.Println(result)

			log.Infof("执行命令“%s”成功", shell)

			return result, err

		}
	}
	err = fmt.Errorf("对话建立失败，执行命令“%s”失败", shell)
	log.Warnf("%v", err)
	return "", err
}

//UploadFile 传输文件模块，输入源地址和目标地址
func (c *Cli) UploadFile(localFilePath string, remotePath string) error {
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		log.Warnf("error:%s", err)

	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)

	dstFile, err := c.SftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join(remotePath, remoteFileName))
		log.Warnf("error:%s", err)

	}
	defer dstFile.Close()

	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localFilePath)
		log.Warnf("error:%s", err)

	}
	dstFile.Write(ff)
	fmt.Println(localFilePath + "  copy file to remote server finished!")
	log.Infof("传输文件“%s”成功", localFilePath)
	return nil
}

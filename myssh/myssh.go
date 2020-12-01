package myssh

import (
	"fmt"
	"net"
	"time"

	"github.com/romberli/log"
	"golang.org/x/crypto/ssh"
)

// Cmd 输入命令，输出结果
func Cmd(cli *Cli, Cmd string) {
	output, err := cli.Run(Cmd)
	if err == nil {
		s
		log.Infof("%s,succeed %v", Cmd, output)
		fmt.Printf("success :%s \n %v", Cmd, output)
	} else {
		log.Warnf("%s,failed!!! error:%v", Cmd, err)
		fmt.Printf("failed :%s \n %v", Cmd, output)
	}
}

//Cli SSH的连接参数结构体
type Cli struct {
	IP         string      //IP地址
	Username   string      //用户名
	Password   string      //密码
	Port       int         //端口号
	client     *ssh.Client //ssh客户端，连接完毕后会生成一个指针
	LastResult string      //最近一次Run的结果，也就是发送指令后返回的结果
}

//NewCli 创建一个新的Cli结构体，指定连接参数，client和LastResult为空
//@param ip IP地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func NewCli(ip string, username string, password string, port int) *Cli {
	cli := new(Cli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}

	return cli
}

//Run 执行shell
//@param shell shell脚本命令
func (c Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

//连接
func (c *Cli) connect() error {
	config := ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	c.client = sshClient
	return nil
}

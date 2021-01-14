package myssh

import (
	"fmt"
	"strconv"

	"github.com/AllinChen/MysqlInstaller/mycfg"
	"github.com/AllinChen/MysqlInstaller/mycnf"
)

//Install 安装流程
func Install(mycfgfile, ip, port string) error {
	//制造flag获得端口号和IP号
	mycfg.UseLog(ip, port)
	//制造conf文件
	mycnf.GenerateMyCnf(ip, port)
	InstallerInfo := mycfg.Read(mycfgfile, "=", ";")
	portused, _ := strconv.Atoi(InstallerInfo["PORT"])
	// fmt.Print(InstallerInfo)
	cli := NewCli(ip, InstallerInfo["USERNAME"], InstallerInfo["PASSWORD"], portused)
	// fmt.Println(cli)
	if err := cli.StartConnect(); err == nil {
		////////还要加上传送文件的过程
		Cfg := mycfg.GetCfg("./src/AutoMysql.cfg")

		cli.UploadFile("./src/mysql.tar.gz", Cfg.MysqlPath+"/")
		cli.Tar(Cfg.MysqlPath + "/mysql.tar.gz")
		cli.Mv("/root/mysql-5.7.31-linux-glibc2.12-x86_64", Cfg.MysqlPath+"/mysql")
		cli.UploadFile("./src/my.cnf", "/etc/")
		//读取配置文件

		//创建用户，用户组
		cli.Useradd("mysql")
		// mkdir -p /data1/mysql3306/binlog
		cli.Mkdir(Cfg.MysqlPath + "/mysql" + port + "/binlog")
		// mkdir -p /data1/mysql3306/data
		cli.Mkdir(Cfg.MysqlPath + "/mysql" + port + "/data")
		// if err != nil{
		// return err
		// }
		// chown -R mysql:mysql /data1/mysql
		cli.Mkdir(Cfg.MysqlPath + "/mysql" + port + "/binlog")

		cli.Chown("mysql:mysql " + Cfg.MysqlPath + "/mysql" + port + "/binlog")

		cli.Chown("mysql:mysql " + Cfg.MysqlPath + "/mysql" + port + "/data")
		cli.Chmod(Cfg.MysqlPath + "/mysql")

		cli.Cp(Cfg.MysqlPath+"/mysql/bin/*", "/usr/bin/")

		// runuser -l mysql -c '/data1/mysql/bin/mysqld --initialize-insecure --user=mysql --basedir=/data1/mysql --datadir=/data1/mysql3306/data'
		// runuser -l mysql -c 'mysqld_multi start 3306'
		cli.Run("runuser -l mysql -c '" + Cfg.MysqlPath + "/mysql/bin/mysqld --initialize-insecure --user=mysql --basedir=" + Cfg.MysqlPath + "/mysql --datadir=" + Cfg.MysqlPath + "/mysql3306/data'")
		return nil
	} else {
		fmt.Println(err)
		return err
	}

}

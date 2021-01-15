package myssh

import (
	"fmt"
	"strconv"

	"github.com/AllinChen/MysqlInstaller/mycfg"
	"github.com/AllinChen/MysqlInstaller/mycnf"
	"github.com/romberli/log"
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
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", cli)
	if err := cli.StartConnect(); err == nil {

		////////还要加上传送文件的过程
		Cfg := mycfg.GetCfg("./src/AutoMysql.cfg")

		_, err = cli.Mkdir(Cfg.MysqlPath)
		if err != nil {
			log.Warnf("建立MYSQL文件夹失败 ")
			return err
		} else {
			log.Infof("建立MYSQL文件夹成功 ")
		}

		err := cli.UploadFile("./src/mysql.tar.gz", Cfg.MysqlPath+"/")
		if err != nil {
			log.Warnf("传输MySQL安装包失败 ")
			return err
		} else {
			log.Infof("传输MySQL安装包成功 ")
		}

		_, err = cli.Tar(Cfg.MysqlPath + "/mysql.tar.gz")
		if err != nil {
			log.Warnf("解压MySQL安装包失败 ")
			return err
		} else {
			log.Infof("解压MySQL安装包成功 ")
		}

		_, err = cli.Mv("/root/mysql-5.7.31-linux-glibc2.12-x86_64", Cfg.MysqlPath+"/mysql")

		err = cli.UploadFile("./src/my.cnf", "/etc/")
		if err != nil {
			log.Warnf("传输cnf文件失败 ")
			return err
		} else {
			log.Infof("传输cnf文件成功 ")
		}
		//读取配置文件

		//创建用户，用户组

		_, err = cli.Useradd("mysql")
		// if err != nil {
		// 	log.Warnf("创建用户失败 ")
		// 	return err
		// } else {
		// 	log.Infof("创建用户成功 ")
		// }
		// mkdir -p /data1/mysql3306/binlog

		_, err = cli.Mkdir(Cfg.MysqlPath + "/mysql" + port + "/binlog")
		if err != nil {
			log.Warnf("建立binlog文件夹失败 ")
			return err
		} else {
			log.Infof("建立binlog文件夹成功 ")
		}
		// mkdir -p /data1/mysql3306/data

		_, err = cli.Mkdir(Cfg.MysqlPath + "/mysql" + port + "/data")
		if err != nil {
			log.Warnf("建立data文件夹失败 ")
			return err
		} else {
			log.Infof("建立data文件夹成功 ")
		}
		// if err != nil{
		// return err
		// }
		// chown -R mysql:mysql /data1/mysql

		_, err = cli.Chown("mysql:mysql " + Cfg.MysqlPath + "/mysql" + port + "/binlog")
		if err != nil {
			log.Warnf("binlog文件夹赋权失败 ")
			return err
		} else {
			log.Infof("binlog文件夹赋权成功 ")
		}

		_, err = cli.Chown("mysql:mysql " + Cfg.MysqlPath + "/mysql" + port + "/data")
		if err != nil {
			log.Warnf("data文件夹赋权失败 ")
			return err
		} else {
			log.Infof("data文件夹赋权成功 ")
		}

		_, err = cli.Chmod(Cfg.MysqlPath + "/mysql")
		if err != nil {
			log.Warnf("文件权限修改失败 ")
			return err
		} else {
			log.Infof("文件权限修改成功 ")
		}

		_, err = cli.Cp(Cfg.MysqlPath+"/mysql/bin/*", "/usr/bin/")
		if err != nil {
			log.Warnf("移动执行文件至系统文件夹失败 ")
			return err
		} else {
			log.Infof("移动执行文件至系统文件夹成功 ")
		}

		// runuser -l mysql -c '/data1/mysql/bin/mysqld --initialize-insecure --user=mysql --basedir=/data1/mysql --datadir=/data1/mysql3306/data'
		// runuser -l mysql -c 'mysqld_multi start 3306'

		_, err = cli.Run("runuser -l mysql -c '" + Cfg.MysqlPath + "/mysql/bin/mysqld --initialize-insecure --user=mysql --basedir=" + Cfg.MysqlPath + "/mysql --datadir=" + Cfg.MysqlPath + "/mysql" + port + "/data'")
		if err != nil {
			log.Warnf("执行MySQL初始化失败 ")
			return err
		} else {
			log.Infof("执行MySQL初始化成功 ")
		}
		log.Infof("MySQL安装成功")
		return nil
	} else {
		fmt.Println(err)
		log.Warnf("安装失败")
		return err
	}

}

package myssh

import (
	"fmt"

	"github.com/AllinChen/MysqlInstaller/conf"
	"github.com/AllinChen/MysqlInstaller/mycfg"
	"github.com/AllinChen/MysqlInstaller/myflag"
)

//Install 安装流程
func Install(c Cli) error {
	//制造flag获得端口号和IP号
	ip, port, _ := myflag.MakeFlag()
	//制造conf文件
	conf.GenerateMyCnf(*ip, *port)

	////////还要加上传送文件的过程
c.UploadFile("./src/my.cnf","/etc/")
	//读取配置文件
	Cfg := mycfg.GetCfg("./src/AutoMysql.cfg")
	// mkdir -p /data1/mysql3306/binlog
	c.Mkdir(Cfg.MysqlPath + "/mysql3306/binlog")
	// mkdir -p /data1/mysql3306/data
	c.Mkdir(Cfg.MysqlPath + "/mysql3306/data")
		// if err != nil{
		// return err	
		// }
	
	// chown -R mysql:mysql /data1/mysql
	c.Mkdir(Cfg.MysqlPath + "/mysql3306/data"))
	// cp -rf ./related/limits.conf /etc/security/limits.conf  64000
	// mkdir -p ./related/
	// mv /data1/mysql731 ./related/mysql
	// cp -rf ./related/mysql/* /data1/mysql

	// cp -rf ./related/my.cnf /etc/my.cnf
	// mkdir -p /data1/mysql3306/data
	// mkdir -p /data1/mysql3306/tmp
	// mkdir -p /data1/mysql3306/sock
	// mkdir -p /data1/mysql3306/log
	// mkdir -p /data1/mysql3306/pid
	// chmod -R 775 /data1/mysql
	// chown -R mysql:mysql /data1/mysql
	// cp -rf /data1/mysql/bin/mysql /usr/bin
	// cp -rf /data1/mysql/bin/mysqld /usr/bin
	// cp -rf /data1/mysql/bin/mysqld_safe /usr/bin
	// cp -rf /data1/mysql/bin/mysqld_multi /usr/bin
	// cp -rf /data1/mysql/bin/mysqldump /usr/bin
	// cp -rf /data1/mysql/bin/mysqlbinlog /usr/bin
	// cp -rf /data1/mysql/bin/mysql_config_editor /usr/bin
	// cp -rf /data1/mysql/bin/my_print_defaults /usr/bin
	// cp -rf /data1/mysql/bin/mysqladmin /usr/bin
	// cp -rf ./related/.bash_profile /home/mysql/.bash_profile
	// runuser -l mysql -c '/data1/mysql/bin/mysqld --initialize-insecure --user=mysql --basedir=/data1/mysql --datadir=/data1/mysql3306/data'
	// runuser -l mysql -c 'mysqld_multi start 3306'
}


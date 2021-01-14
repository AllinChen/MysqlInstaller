package main

import (
	"github.com/AllinChen/MysqlInstaller/myflag"
	"github.com/AllinChen/MysqlInstaller/mygin"
	"github.com/AllinChen/MysqlInstaller/myssh"
)

func main() {
	const (
		mycfgfile = "./src/AutoMysql.cfg"
	)

	ip, port, _ := myflag.MakeFlag()
	if *ip == "0.0.0.0" {
		//*ip = "192.168.171.141"
		mygin.StartGin(mycfgfile, ip, port)

	} else {
		myssh.Install(mycfgfile, *ip, *port)
	}

	// ip := "192.168.171.148"
	// InstallerInfo := mycfg.Read(mycfgfile, "=", ";")
	// portused, _ := strconv.Atoi(InstallerInfo["PORT"])
	// // fmt.Print(InstallerInfo)
	// cli := myssh.NewCli(ip, InstallerInfo["USERNAME"], InstallerInfo["PASSWORD"], portused)
	// fmt.Println(cli)
	// if err := cli.StartConnect(); err == nil {
	// 	//cli.Run("cat /etc/my.cnf")

	// }
}

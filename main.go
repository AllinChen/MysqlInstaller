package main

import (
	"github.com/AllinChen/MysqlInstaller/myflag"
	"github.com/AllinChen/MysqlInstaller/myssh"
)

func main() {
	const (
		mycfgfile = "./src/AutoMysql.cfg"
	)
	// InstallerInfo := mycfg.Read(mycfgfile, "=", ";")
	// port, _ := strconv.Atoi(InstallerInfo["PORT"])
	// fmt.Print(InstallerInfo)
	// cli := myssh.NewCli(InstallerInfo["IP"], InstallerInfo["USERNAME"], InstallerInfo["PASSWORD"], port)
	// if err := cli.StartConnect(); err == nil {
	// 	cli.UploadFile("./src/my.cnf", "/etc/")
	// } else {
	// 	fmt.Println(err)
	// }
	ip, port, _ := myflag.MakeFlag()
	if *ip == "0.0.0.0" {
		*ip = "192.168.171.141"

	}
	myssh.Install(mycfgfile, *ip, *port)

}

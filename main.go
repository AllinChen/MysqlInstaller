package main

import (
	"fmt"
	"strconv"

	"github.com/AllinChen/AutoMycnf/mycfg"
	"github.com/AllinChen/MysqlInstaller/myssh"
)

func main() {
	const (
		mycfgfile = "./config/Installer.cfg"
	)
	InstallerInfo := mycfg.Read(mycfgfile, ":", "\n")
	port, _ := strconv.Atoi(InstallerInfo["Port"])
	fmt.Print(InstallerInfo)
	cli := myssh.NewCli(InstallerInfo["Ip"], InstallerInfo["User"], InstallerInfo["Password"], port)
	if err := cli.StartConnect(); err == nil {
		cli.Run("mysql.server start")
		cli.Run("df -h")
		cli.Run("top")
	} else {
		fmt.Println(err)
	}

}

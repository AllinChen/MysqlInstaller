package pg

import "github.com/AllinChen/MyCfg/mycfg"

func Install(cfgname string) {
	mycfg.Read(cfgname)
}

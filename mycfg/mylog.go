package mycfg

import (
	"fmt"
	"os"
	"runtime"

	"github.com/romberli/log"
)

//UseLog 建立日志
func UseLog(ip, port string) error {
	dir, _ := os.Getwd()
	var logname string
	if runtime.GOOS == "windows" {
		logname = dir + "\\logs\\" + ip + "_" + port + ".log"
	} else {
		logname = dir + "/logs/" + ip + "_" + port + ".log"
	}
	_, _, err := log.InitLoggerWithDefaultConfig(logname)
	if err != nil {
		fmt.Println("init log failed")
		return fmt.Errorf("init log failed")
	}
	return nil
}

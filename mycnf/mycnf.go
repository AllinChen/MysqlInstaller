package mycnf

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/AllinChen/AutoMysql/mycfg"
	"github.com/AllinChen/MysqlInstaler/myssh"
	"github.com/romberli/log"
)

//ReplaceCnf 对配置文件中的值进行修改
func ReplaceCnf(cnfname, cnfoptionname, datatoreplace string) (NewClicnf string) {

	value, project := ReGet(cnfoptionname, cnfname)
	log.Infof("替换 %s,原值 %s,更改为 %s", cnfoptionname, value, datatoreplace)
	NewClicnf = strings.Replace(cnfname, project, cnfoptionname+"="+datatoreplace, -1)
	return

}

//myssh mycfg
//StringRead 配置文件读取模块
func StringRead(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Warnf("读取文件失败:%#v", err)
		return ""
	}
	log.Info("读取文件成功")
	return string(f)
}

//ReGet 获取目标项
func ReGet(compile, cfgfile string) (res, respro string) {
	re := regexp.MustCompile(compile + `=([^\n]+)`)
	//解释规则，解析正则表达式，如果成功则返回
	if re == nil {
		fmt.Println("error regexp")
		panic(re)
	}
	//根据规则提取信息

	result := re.FindAllStringSubmatch(cfgfile, -1)
	//fmt.Println(result)
	if result == nil {

		log.Warnf("get %s failed", compile)
		return " ", " "
	}
	fmt.Println(compile, " = ", result[0][1])
	res = result[0][1]
	log.Infof("get %s success , value = %s", compile, res)
	respro = compile + "=" + res
	return

}

//ReGetAll 获取目标项
func ReGetAll(cfgfile string) (res [][]string) {
	re := regexp.MustCompile(`([^\n]+)=([^\n]+)`)
	//解释规则，解析正则表达式，如果成功则返回
	if re == nil {
		fmt.Println("error regexp")
		panic(re)
	}
	//根据规则提取信息

	res = re.FindAllStringSubmatch(cfgfile, -1)
	//fmt.Println(result)

	return

}

//ChangeCnf 自动修改cnf文件
func ChangeCnf(cli *myssh.Cli, cfg *mycfg.Cfg) string {
	mycnftxt := StringRead(cfg.LocalMycnfPath)

	EzGet("key_buffer_size", mycnftxt) //默认8M 大于4G设置256M  free -h
	EzGet("sort_buffer_size", mycnftxt)
	EzGet("thread_cache_size", mycnftxt)       //1G内存设置为8；2G设置为16
	EzGet("innodb_buffer_pool_size", mycnftxt) //物理内存的80%
	EzGet("innodb_log_buffer_size", mycnftxt)
	EzGet("innodb_log_file_size", mycnftxt)

	var Memory int
	if Mem, err := cli.Run("free| grep Mem"); err == nil {
		fmt.Println(strings.Replace(Mem, " ", "/", -1))
		Memory = ReGetMem(Mem) / 1024
	}
	//为了使单位统一，获取的内存数值为byte单位
	//需要进行转换得到M和G为单位的内存大小
	Mem := Memory / 1024
	if Mem >= 1 && Mem < 4 {
		mycnftxt = ReplaceCnf(mycnftxt, "key_buffer_size", "8M")
		mycnftxt = ReplaceCnf(mycnftxt, "thread_cache_size", fmt.Sprint(Mem*8))
		mycnftxt = ReplaceCnf(mycnftxt, "innodb_buffer_pool_size", fmt.Sprint(Mem*800)+"M")
	} else if Mem >= 4 {
		mycnftxt = ReplaceCnf(mycnftxt, "key_buffer_size", "256M")
		mycnftxt = ReplaceCnf(mycnftxt, "thread_cache_size", fmt.Sprint(Mem*8))
		mycnftxt = ReplaceCnf(mycnftxt, "innodb_buffer_pool_size", fmt.Sprint(Mem*800)+"M")
	} else {
		mycnftxt = ReplaceCnf(mycnftxt, "key_buffer_size", "8M")
		mycnftxt = ReplaceCnf(mycnftxt, "thread_cache_size", "8")
		mycnftxt = ReplaceCnf(mycnftxt, "innodb_buffer_pool_size", "500M")

	}

	mycnftxt = ReplaceCnf(mycnftxt, "socket", cfg.MysqlPath+"/data/mysql.sock")
	mycnftxt = ReplaceCnf(mycnftxt, "basedir", cfg.MysqlPath+"/mysql731")
	mycnftxt = ReplaceCnf(mycnftxt, "datadir", cfg.MysqlPath+"/data")
	mycnftxt = ReplaceCnf(mycnftxt, "log_error", cfg.MysqlPath+"/data/mysql.err")
	mycnftxt = ReplaceCnf(mycnftxt, "pid-file", cfg.MysqlPath+"/data/mysql.pid")
	mycnftxt = ReplaceCnf(mycnftxt, "log-bin", cfg.MysqlPath+"/binlog/mysql-bin")
	mycnftxt = ReplaceCnf(mycnftxt, "relay_log", cfg.MysqlPath+"/binlog/relay-bin")
	mycnftxt = ReplaceCnf(mycnftxt, "report_host", cfg.IP)
	mycnftxt = ReplaceCnf(mycnftxt, "slow_query_log_file", cfg.MysqlPath+"/data/slow_query.log")

	// res := ReGetAll(mycnftxt)
	// for _, options := range res {
	// 	fmt.Println(options[1], " = ", options[2])
	// }

	// 保存到文件
	return mycnftxt

}

//AutoCnf 自动根据配置更改cnf文件，单机
func AutoCnf(cfgName string) {
	dir, _ := os.Getwd()
	cfg := mycfg.GetCfg(dir + "\\" + cfgName)
	cli := myssh.NewCli(cfg.IP, cfg.Username, cfg.Password, cfg.Port)
	mycnftxt := ChangeCnf(cli, cfg)
	str := []byte(mycnftxt)
	ioutil.WriteFile(`my.cnf`, str, 0666)

}

//AutoCnfM 自动根据配置更改cnf文件，主机
func AutoCnfM(cfgName string) {
	dir, _ := os.Getwd()
	cfg := mycfg.GetCfg(dir + "\\" + cfgName)
	cli := myssh.NewCli(cfg.IP, cfg.Username, cfg.Password, cfg.Port)
	mycnftxt := ChangeCnf(cli, cfg)
	mycnftxt = ReplaceCnf(mycnftxt, "server-id", "1")
	str := []byte(mycnftxt)
	ioutil.WriteFile(`my.cnf`, str, 0666)

}

//AutoCnfS 自动根据配置更改cnf文件，从机
func AutoCnfS(cfgName string) {
	dir, _ := os.Getwd()
	cfg := mycfg.GetCfg(dir + "\\" + cfgName)
	cli := myssh.NewCli(cfg.IP, cfg.Username, cfg.Password, cfg.Port)
	mycnftxt := ChangeCnf(cli, cfg)
	mycnftxt = ReplaceCnf(mycnftxt, "server-id", "2")
	str := []byte(mycnftxt)
	ioutil.WriteFile(`my.cnf`, str, 0666)

}

//EzGet 快速得到对应项目的值
func EzGet(name, filetxt string) {
	ReGet(name, filetxt)
	return
}

//ReGetMem 得到内存
func ReGetMem(inf string) (res int) {
	re := regexp.MustCompile(`//([0-9A-Z]+)//`)
	//解释规则，解析正则表达式，如果成功则返回
	if re == nil {
		fmt.Println("error regexp")
		panic(re)
	}
	result := re.FindAllStringSubmatch(strings.Replace(inf, " ", "/", -1), -1)
	//fmt.Println(result)
	if result == nil {

		log.Warnf("get mem failed")
		return 0
	}
	//fmt.Println(result)
	res, _ = strconv.Atoi(result[0][1])

	log.Infof("get mem success , value = %s", res)

	return
}

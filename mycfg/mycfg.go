package mycfg

import (
	"fmt"

	"regexp"
	"strconv"

	"github.com/romberli/log"
)

//Cfg 配置结构体
type Cfg struct {
	IP string
	//IP地址
	Username string
	//用户名
	Password string
	//密码
	Port int
	//端口号
	LocalMycnfPath string
	//本地mycnf路径
	RemoteCnfPath string
	//服务器mycnf路径
	MysqlPath string
	//指定MySQL安装路径
	InstallSQLPath string
	//指定初始化SQL文件位置
	MysqlTarPath string
	//指定mysql安装包位置
}

//GetCfg 得到配置信息
func GetCfg(file string) (cfg *Cfg) {
	cfgfile := StringRead(file)
	IP := ReGet("IP", cfgfile)
	Username := ReGet("USERNAME", cfgfile)
	PassWord := ReGet("PASSWORD", cfgfile)
	Port := ReGet("PORT", cfgfile)
	RealPort, _ := strconv.Atoi(Port)
	LocalMycnfPath := ReGet("LOCALMYCNFPATH", cfgfile)
	RemoteCnfPath := ReGet("REMOTECNFPATH", cfgfile)
	MysqlPath := ReGet("MYSQLPATH", cfgfile)
	InstallSQLPath := ReGet("INSTALLSQLPATH", cfgfile)
	MysqlTarPath := ReGet("MYSQLTARPATH", cfgfile)
	cfg = &Cfg{IP, Username, PassWord, RealPort, LocalMycnfPath, RemoteCnfPath, MysqlPath, InstallSQLPath, MysqlTarPath}
	if cfg.IP == "" || cfg.Username == "" || cfg.Password == "" || cfg.LocalMycnfPath == "" || cfg.RemoteCnfPath == "" || cfg.MysqlPath == "" || cfg.InstallSQLPath == "" || MysqlTarPath == "" {
		log.Warn("配置表建立失败，需检查配置文件是否正确建立")
		return nil
	}
	log.Info("成功建立CFG配置信息")
	return cfg

}

//ReGet 配置单独项的匹配模块
func ReGet(compile, cfgfile string) (res string) {
	re := regexp.MustCompile(compile + `=([^;]+)`)

	//解释规则，解析正则表达式，如果成功则返回
	if re == nil {
		fmt.Println("error regexp")
		panic(re)
	}

	//根据规则提取信息
	result := re.FindAllStringSubmatch(cfgfile, -1)
	if result == nil {

		//	log.Warnf("get %s failed", compile)
		return " "
	}
	fmt.Println(compile, " = ", result[0][1])
	res = result[0][1]

	//log.Infof("get %s success , value = %s", compile, res)
	return

}

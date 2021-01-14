package mygin

import (
	"fmt"
	"regexp"

	"github.com/AllinChen/MysqlInstaller/myssh"
	"github.com/gin-gonic/gin"
	"github.com/romberli/log"
)

//StartGin 展开Gin服务
func StartGin(mycfgfile string, ip, port *string) (err error) {
	r := gin.Default()

	r.GET("/installmysql/:ipport", func(c *gin.Context) {
		ipport := c.Param("ipport")
		*ip, *port, _ = Reipport(ipport)
		c.String(200, "ip:  %s\nport:  %s", *ip, *port)

		myssh.Install(mycfgfile, *ip, *port)

	})
	r.Run(":8080")
	return nil
}

//Reipport 解析ip和port
func Reipport(ipport string) (ip, port string, err error) {
	re := regexp.MustCompile(`([^:]+):([0-9]+)`)
	if re == nil {
		fmt.Println("error regexp")
		//panic(re)
		return "", "", fmt.Errorf("error regexp")
	}
	//根据规则提取信息
	result := re.FindAllStringSubmatch(ipport, -1)
	if result == nil {

		log.Warnf("解析URL中的IP地址失败 %s", ipport)
		return "", "", fmt.Errorf("解析失败 %s", ipport)
	}
	fmt.Println(result[0][0])
	return result[0][1], result[0][2], nil
}

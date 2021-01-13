package myflag

import (
	"flag"
	"fmt"
)

// func main() {

// 	ip, port, _ := MakeFlag()
// 	fmt.Println(*ip, "port has value ", *port)
// }

//MakeFlag 制造flag,导入IP地址和端口号，端口号默认3306
func MakeFlag() (ip, port *string, Error error) {

	ip = flag.String("ip", "0.0.0.0", "ip address")
	port = flag.String("port", "3306", "端口号，默认为3306")
	flag.Parse()
	if *ip == "" {
		Error = fmt.Errorf("NILL IP,ERROR")
		return ip, port, Error
	}
	// fmt.Println("ip has value ", *ip)
	// fmt.Println("port has value ", *port)
	return ip, port, nil
}

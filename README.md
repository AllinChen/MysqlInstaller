# MysqlInstaller

功能介绍
    通过不同方法输入IP地址和端口号，实现自动安装MySQL

模块介绍
    1、mycfg:简单的配置文件解析，读取配置文件中单项配置的信息
    2、mycnf:根据配置文件和输入的IP和端口信息，自动组装my.cnf文件
    3、myflag:实现通过flag的方法输入参数的功能
    4、mygin:实现了通过gin框架输入参数的功能
    5、myssh:实现了通过ssh远程执行命令和传输文件的功能
    6、logs:存放日志文件，文件格式 ip地址_端口.log
    7、src:存放配置文件，my.cnf,和MySQL安装包(推荐MySQL7.31,并将安装包命名为mysql.tar.gz)

使用介绍
    1、通过flag的方法输入参数
    `go run main.go -ip "192.168.171.149" -port "3306"`
    2、通过网页方式输入参数
    运行main.go
    在浏览器地址栏中输入 `localhost:8080\installmysql\192.168.171.149:3306`

配置文件参数介绍
   `IP             string //IP地址`
   `Username       string //用户名`
   `Password       string //密码`
   `Port           int    //端口号`
   `LocalMycnfPath string //本地mycnf路径`
   `RemoteCnfPath  string //服务器mycnf路径`
   `MysqlPath      string //指定MySQL安装路径`
   `InstallSQLPath string //指定初始化SQL文件位置`
   `MysqlTarPath   string //指定mysql安装包位置`
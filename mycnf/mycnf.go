package mycnf

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/AllinChen/MysqlInstaller/mycfg"
	"github.com/BurntSushi/toml"
	"github.com/romberli/log"
)

//订制配置文件解析载体
type Config struct {
	Client *Client
	Mysql  *Mysql
	Mysqld *Mysqld
}

//订制Client语句结构
type Client struct {
	Socket string `toml:"socket"`
}
type Mysql struct {
	DefaultCharacterSet string `toml:"default-character-set"`
}
type Mysqld struct {
	Mysqld        string `toml:"mysqld"`
	Port          string `toml:"port"`
	LogTimeStamps string `toml:"log_timestamps"`
	BaseDir       string `toml:"basedir"`
	DataDir       string `toml:"datadir"`
	TmpDir        string `toml:"tmpdir"`
	Socket        string `toml:"socket"`
	LogError      string `toml:"log_error"`
	PidFile       string `toml:"pid_file"`

	SeverId string `toml:"server-id"`

	TransactionIsolation string `toml:"transaction-isolation"`
	CharacterSetServer   string `toml:"character_set_server"`
	OpenFilesLimit       string `toml:"open_files_limit"`
	LowerCaseTableNames  string `toml:"lower_case_table_names"`
	MaxConnections       string `toml:"max_connections"`
	MaxConnectErrors     string `toml:"max_connect_errors"`
	ConnectTimeout       string `toml:"connect_timeout"`
	LockWaitTimeout      string `toml:"lock_wait_timeout"`
	WaitTimeout          string `toml:"wait_timeout"`
	InteractiveTimeout   string `toml:"interactive_timeout"`
	MaxAllowedPacket     string `toml:"max_allowed_packet"`
	ThreadCacheSize      string `toml:"thread_cache_size"`

	SqlMode string `toml:"sql_mode"`

	BinlogFormat   string `toml:"binlog_format"`
	BinlogRowImage string `toml:"binlog_row_image"`

	LogBin            string `toml:"log-bin"`
	MaxBinlogSize     string `toml:"max_binlog_size"`
	ExpireLogsDays    string `toml:"expire_logs_days"`
	BinlogErrorAction string `toml:"binlog_error_action"`

	LogSlaveUpdates        string `toml:"log_slave_updates"`
	RelayLog               string `toml:"relay_log"`
	MaxRelayLogSize        string `toml:"max_relay_log_size"`
	RelayLogPurge          string `toml:"relay_log_purge"`
	RelayLogRecovery       string `toml:"relay_log_recovery"`
	MasterInfoRepository   string `toml:"master_info_repository"`
	RelayLogInfoRepository string `toml:"relay_log_info_repository"`

	ReportHost string `toml:"report_host"`
	ReportPort string `toml:"report_port"`

	SyncBinlog string `toml:"sync_binlog"`

	InnodbFlushLogAtTrxCommit        string `toml:"innodb_flush_log_at_trx_commit"`
	InnodbBufferPoolSize             string `toml:"innodb_buffer_pool_size"`
	InnodbSortBufferSize             string `toml:"innodb_sort_buffer_size"`
	InnodbLogBufferSize              string `toml:"innodb_log_buffer_size"`
	InnodbLogFileSize                string `toml:"innodb_log_file_size"`
	InnodbLogFilesInGroup            string `toml:"innodb_log_files_in_group"`
	InnodbLockWaitTimeout            string `toml:"innodb_lock_wait_timeout"`
	InnodbLockGroupHomeDir           string `toml:"innodb_lock_group_home_dir"`
	InnodbIoCapacity                 string `toml:"innodb_io_capacity"`
	InnodbIoCapacityMax              string `toml:"innodb_io_capacity_max"`
	InnodbFilePerTable               string `toml:"innodb_file_per_table"`
	InnodbStatsPersistentSamplePages string `toml:"innodb_stats_persistent_sample_pages"`
	InnodbOnlineAlterLogMaxSize      string `toml:"innodb_online_alter_log_max_size"`
	InnodbThreadConcurrency          string `toml:"innodb_thread_concurrency"`
	InnodbWriteIoThreads             string `toml:"innodb_write_io_threads"`
	InnodbReadIoThreads              string `toml:"innodb_read_io_threads"`
	InnodbPageCleaners               string `toml:"innodb_page_cleaners"`
	InnodbFlushMethod                string `toml:"innodb_flush_method"`

	InnodbMonitorEnable     string `toml:"innodb_monitor_enable"`
	InnodbPrintAllDeadlocks string `toml:"innodb_print_all_deadlocks"`

	GtidMode                 string `toml:"gitd_mode"`
	EnforceGtidConsistency   string `toml:"enforce_gtid_consistency"`
	BinlogGtidSimpleRecovery string `toml:"binlog_gtid_simple_recovery"`
	SlaveParallelType        string `toml:"slave-parallel-type"`
	SlaveParallelWorkers     string `toml:"slave-parallel-workers"`
	SlavePreserveCommitOrder string `toml:"slave_preserve_commit_order"`
	SlaveTransactionRetries  string `toml:"slave_transaction_retries"`

	LooseInnodbNumaInterLeave        string `toml:"loose_innodb_numa_interleave"`
	InnodbBufferPoolDumpPct          string `toml:"innodb_buffer_pool_dump_pct"`
	InnodbUndoDirectory              string `toml:"innodb_undo_directory"`
	InnodbUndoLogs                   string `toml:"innodb_undo_logs"`
	InnodbUndoTablespaces            string `toml:"innodb_undo_tablespaces"`
	InnodbUndoLogTruncate            string `toml:"innodb_undo_log_truncate"`
	InnodbMaxUndoLogSize             string `toml:"innodb_max_undo_log_size"`
	InnodbPurgeRsegTruncateFrequency string `innodb_purge_rseg_truncate_frequency`

	TableOpenCache       string `toml:"table_open_cache"`
	TmpTableSize         string `toml:"tmp_table_size"`
	MaxHeapTableSize     string `toml:"max_heap_table_size"`
	SortBufferSize       string `toml:"sort_buffer_size"`
	JoinBufferSize       string `toml:"join_buffer_size"`
	ReadBufferSize       string `toml:"read_buffer_size"`
	ReadRndBufferSize    string `toml:"read_rnd_buffer_size"`
	KeyBufferSize        string `toml:"key_buffer_size"`
	BulkInsertBufferSize string `toml:"bulk_insert_buffer_size"`
	BinlogCacheSize      string `toml:"binlog_cache_size"`

	SlowQueryLogFile                  string `toml:"slow_query_log_file"`
	SlowQueryLog                      string `toml:"slow_query_log"`
	LongQueryTime                     string `toml:"long_query_time"`
	LogOutput                         string `toml:"log_output"`
	LogSlowAdminStatements            string `toml:"log_slow_admin_statements"`
	LogSlaveSlaveStatements           string `toml:"log_slave_slave_statements"`
	LogThrottleQueriesNotUsingIndexes string `toml:"log_throttle_queries_not_using_indexes"`

	PerformanceSchema           string `toml:"performance_schema"`
	PerFormanceSchemaInstrument string `toml:"performance-schema-instrument"`

	PluginLoadAdd          string `toml:"plugin-load-add"`
	ValidatePassword       string `toml:"validate-password"`
	ValidatePasswordPolicy string `toml:"validate_password_policy"`

	SymbolicLinks string `toml:"symbolic-links"`
}

var Conf *Config = new(Config)

func init() {
	//读取配置文件
	_, err := toml.DecodeFile("./mycnf/test.cnf", Conf)
	if err != nil {
		fmt.Println(err)
	}
}
func ModifyMysqld(attribute, val string) *Config {
	v := reflect.ValueOf(Conf.Mysqld).Elem()
	v.FieldByName(attribute).Set(reflect.ValueOf(val))
	return Conf
}

// 结构体数据转化为字符串
func GetBuffer(Struct interface{}) string {
	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	encoder.Encode(Struct)
	fmt.Println(buffer.String())
	return buffer.String()
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func GenerateMyCnf(Ip string, PortNum string) error {
	filename := "./src/my.cnf"
	var f *os.File
	var err error
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, os.ModeAppend) //打开文件
		fmt.Println("文件存在")

	} else {
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
		log.Warnf("无法找到默认cnf配置")
	}
	if err != nil {
		return err
	}

	// 开始拼接字符串
	// client section
	result := "[client]\n"

	Conf.Client = GenerateClient(Conf.Client, PortNum)
	result = result + GetBuffer(Conf.Client) + "\n"

	// mysql section
	result = result + "[mysql]\n"
	Conf.Mysql = GenerateMysql(Conf.Mysql)
	result = result + GetBuffer(Conf.Mysql)
	result = result + "auto-rehash\n\n"

	// mysqld section
	Conf.Mysqld = GenerateMysqld(Conf.Mysqld, Ip, PortNum)
	result = result + "[mysqld" + PortNum + "]\n"
	result = result + GetBuffer(Conf.Mysqld)
	result += "secure_file_priv=\nskip-host-cache\nskip-name-resolve"

	// 开始写入
	_, err = io.WriteString(f, result) //写入文件(字符串)
	if err != nil {
		return err
	}
	log.Infof("my.cnf建立成功")
	return nil
}

// 给Client结构体赋值
func GenerateClient(client *Client, PortNum string) *Client {
	Cfg := mycfg.GetCfg("./src/AutoMysql.cfg")
	client.Socket = Cfg.MysqlPath + "/mysql" + PortNum + "/mysql.sock"
	fmt.Println(client.Socket)
	log.Infof("client模块创建成功")
	return client
}

// 给Mysql结构体赋值
func GenerateMysql(mysql *Mysql) *Mysql {
	mysql.DefaultCharacterSet = "utf8mb4"
	log.Infof("mysql模块创建成功")
	return mysql
}

// 给Mysqld结构体赋值
func GenerateMysqld(mysqld *Mysqld, Ip string, PortNum string) *Mysqld {
	// secure_file_priv=
	// skip-host-cache
	// skip-name-resolve
	Cfg := mycfg.GetCfg("./src/AutoMysql.cfg")
	mysqld.Mysqld = "/usr/bin/msqld"
	mysqld.Port = PortNum
	mysqld.LogTimeStamps = "SYSTEM"
	mysqld.BaseDir = Cfg.MysqlPath + "/mysql"

	mysqld.DataDir = Cfg.MysqlPath + "/mysql" + PortNum + "/data"
	mysqld.TmpDir = Cfg.MysqlPath + "/mysql" + PortNum + "/data"
	mysqld.Socket = Cfg.MysqlPath + "/mysql" + PortNum + "/mysql.sock"
	mysqld.LogError = Cfg.MysqlPath + "/mysql" + PortNum + "/data/mysqld.err"
	mysqld.PidFile = Cfg.MysqlPath + "/mysql" + PortNum + "/mysql.pid"

	mysqld.SeverId = PortNum + strings.Split(Ip, ".")[3]
	mysqld.TransactionIsolation = "READ-COMMITTED"
	mysqld.CharacterSetServer = "utf8mb4"
	mysqld.OpenFilesLimit = "65535"
	mysqld.LowerCaseTableNames = "1"
	mysqld.MaxConnections = "200"
	mysqld.MaxConnectErrors = "100000000"
	mysqld.ConnectTimeout = "10"
	mysqld.LockWaitTimeout = "3600"
	mysqld.WaitTimeout = "86400"
	mysqld.InteractiveTimeout = "86400"
	mysqld.MaxAllowedPacket = "64M"
	mysqld.ThreadCacheSize = "512"

	mysqld.SqlMode = "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,NO_AUTO_VALUE_ON_ZERO"

	mysqld.BinlogFormat = "row"
	mysqld.BinlogRowImage = "full"

	mysqld.LogBin = Cfg.MysqlPath + "/mysql" + PortNum + "/binlog/mysql-bin"

	mysqld.MaxBinlogSize = "1G"
	mysqld.ExpireLogsDays = "7"
	mysqld.BinlogErrorAction = "ABORT_SERVER"

	mysqld.LogSlaveUpdates = "1"

	mysqld.RelayLog = Cfg.MysqlPath + "/mysql" + PortNum + "/binlog/mysql-relay"
	mysqld.MaxRelayLogSize = "1G"
	mysqld.RelayLogPurge = "0"
	mysqld.RelayLogRecovery = "1"
	mysqld.MasterInfoRepository = "TABLE"
	mysqld.RelayLogInfoRepository = "TABLE"

	mysqld.ReportHost = Ip
	mysqld.ReportPort = PortNum

	mysqld.SyncBinlog = "1"
	mysqld.InnodbFlushLogAtTrxCommit = "1"
	mysqld.InnodbBufferPoolSize = "1024M"
	mysqld.InnodbSortBufferSize = "4M"
	mysqld.InnodbLogBufferSize = "32M"
	mysqld.InnodbLogFileSize = "1G"
	mysqld.InnodbLogFilesInGroup = "4"
	mysqld.InnodbLockWaitTimeout = "60"
	mysqld.InnodbLockGroupHomeDir = Cfg.MysqlPath + "/mysql" + PortNum + "/data"

	mysqld.InnodbIoCapacity = "1000"
	mysqld.InnodbIoCapacityMax = "2000"
	mysqld.InnodbFilePerTable = "1"
	mysqld.InnodbStatsPersistentSamplePages = "64"
	mysqld.InnodbOnlineAlterLogMaxSize = "1G"
	mysqld.InnodbThreadConcurrency = "0"
	mysqld.InnodbWriteIoThreads = "16"
	mysqld.InnodbReadIoThreads = "16"
	mysqld.InnodbPageCleaners = "16"
	mysqld.InnodbFlushMethod = "O_DIRECT"

	mysqld.InnodbMonitorEnable = "all"
	mysqld.InnodbPrintAllDeadlocks = "1"

	mysqld.GtidMode = "on"
	mysqld.EnforceGtidConsistency = "1"
	mysqld.BinlogGtidSimpleRecovery = "1"
	mysqld.SlaveParallelType = "LOGICAL_CLOCK"
	mysqld.SlaveParallelWorkers = "16"
	mysqld.SlavePreserveCommitOrder = "1"
	mysqld.SlaveTransactionRetries = "128"

	mysqld.LooseInnodbNumaInterLeave = "1"
	mysqld.InnodbBufferPoolDumpPct = "40"
	mysqld.InnodbUndoDirectory = Cfg.MysqlPath + "/mysql" + PortNum + "/binlog/"
	mysqld.InnodbUndoLogs = "128"
	mysqld.InnodbUndoTablespaces = "0"
	mysqld.InnodbUndoLogTruncate = "1"
	mysqld.InnodbMaxUndoLogSize = "2G"
	mysqld.InnodbPurgeRsegTruncateFrequency = "128"

	mysqld.TableOpenCache = "2048"
	mysqld.TmpTableSize = "64M"
	mysqld.MaxHeapTableSize = "64M"
	mysqld.SortBufferSize = "4M"
	mysqld.JoinBufferSize = "4M"
	mysqld.ReadBufferSize = "8M"
	mysqld.ReadRndBufferSize = "4M"
	mysqld.KeyBufferSize = "32M"
	mysqld.BulkInsertBufferSize = "64M"
	mysqld.BinlogCacheSize = "1M"

	mysqld.SlowQueryLogFile = Cfg.MysqlPath + "/mysql" + PortNum + "/data/mysql-slow.log"
	mysqld.SlowQueryLog = "ON"
	mysqld.LongQueryTime = "1"
	mysqld.LogOutput = "file"
	mysqld.LogSlowAdminStatements = "1"
	mysqld.LogSlaveSlaveStatements = "1"
	mysqld.LogThrottleQueriesNotUsingIndexes = "10"

	mysqld.PerformanceSchema = "ON"
	mysqld.PerFormanceSchemaInstrument = "wait/lock/metadata/sql/dml/memory=ON"

	mysqld.PluginLoadAdd = "validate_password.so"
	mysqld.ValidatePassword = "FORCE_PLUS_PERMANENT"
	mysqld.ValidatePasswordPolicy = "MEDIUM"
	mysqld.SymbolicLinks = "0"
	log.Infof("mysqld模块创建成功")
	return mysqld

}

/**
 * @Author Mr.LiuQH
 * @Description 数据库相关的配置
 * @Date 2021/2/4 4:14 下午
 **/
package config

/**
 * @description: mysql配置
 * @user: Mr.LiuQH
 */
type MysqlConfig struct {
	Host        string `ini:"host"`
	Port        string `ini:"port"`
	Database    string `ini:"database"`
	UserName    string `ini:"userName"`
	Password    string `ini:"password"`
	Charset     string `ini:"charset"`
	MaxIdleConn int    `ini:"max_idle_conn"`
	MaxOpenConn int    `ini:"max_open_conn"`
	ParseTime   string `ini:"parse_time"`
	Loc         string `ini:"loc"`
	Timeout     string `ini:"timeout"`
	MaxLifeTime string `ini:"max_life_time"`
	TablePre    string `ini:"table_pre"`
	SlowSqlTime string `ini:"slow_sql_time"`
	PrintSqlLog bool   `ini:"print_sql_log"`
}

/**
 * @description: Redis配置信息
 * @user: Mr.LiuQH
 */
type RedisConfig struct {
	Host      string `ini:"host"`
	Port      string `ini:"port"`
	DefaultDB int    `ini:"default_db"`
	Password  string `ini:"password"`
	PoolSize  int    `ini:"pool_size"`
	TimeOut   string `ini:"timeout"`
}

/**
 * @description: 日志配置
 * @user: Mr.LiuQH
 */
type LogrusConfig struct {
	Path         string `ini:"path"`
	Level        string `ini:"level"`
	Formatter    string `ini:"formatter"`
	OutputType   string `ini:"output_type"`
	ReportCaller bool   `ini:"report_caller"`
	ClassSaveNum uint   `ini:"class_save_num"`
	Suffix       string `ini:"suffix_format"`
	SplitTime    string `ini:"split_time"`
}

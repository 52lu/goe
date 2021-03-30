package app

import (
	"fmt"
	"github.com/go-ini/ini"
	. "goe/app/common"
	_ "goe/app/controllers/v1"
	_ "goe/app/controllers/v2"
	"net/http"
)

type App struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	Env  string `ini:"env"`
}

var err error

/**
 * @description: 启动服务
 * @user: Mr.LiuQH
 * @receiver app App
 * @date 2021-02-03 10:21:20
 */
func (app *App) Start() {
	// 捕获启动时错误
	defer BusErrorInstance.CatchError()
	// 加载配置文件
	app.loadConfig()
	// 初始化语句
	app.initializeHandle()
	// 设置打印的信息
	CliInfoInstance.Host = app.Host
	CliInfoInstance.Port = app.Port
	CliInfoInstance.PrintRunMsg()
	// 启动服务
	err := http.ListenAndServe(app.Host+":"+app.Port, RouteListInstance)
	if err != nil {
		fmt.Println("启动失败: " + err.Error())
	}
}

/**
 * @description: 加载配置文件
 * @user: Mr.LiuQH
 * @receiver app
 * @date 2021-02-04 15:29:17
 */
func (app *App) loadConfig() {
	iniPath := ConfigPath + app.Env + ".ini"
	CliInfoInstance.ConfigFile = iniPath
	cfg, err := ini.Load(iniPath)
	BusErrorInstance.ThrowError(err)
	err = cfg.Section("app").MapTo(app)
	// 加载mysql配置
	err = cfg.Section("mysql").MapTo(MysqlConfigInstance)
	// 加载redis配置
	err = cfg.Section("redis").MapTo(RedisConfigInstance)
	// 加载日志配置
	err = cfg.Section("log").MapTo(LogrusConfigInstance)
	// 加载ES配置
	err = cfg.Section("elastic").MapTo(ElasticConfigInstance)
	BusErrorInstance.ThrowError(err)
}

/**
 * @description: 设置数据库连接
 * @user: Mr.LiuQH
 * @receiver app
 * @date 2021-02-04 17:06:29
 */
func (app *App) initializeHandle() {
	fmt.Println(MysqlConfigInstance)
	fmt.Println(RedisConfigInstance)
	fmt.Println(LogrusConfigInstance)
	fmt.Println(ElasticConfigInstance)
	if MysqlConfigInstance.Enabled {
		// 连接Mysql
		connectMysql()
	}
	if RedisConfigInstance.Enabled {
		// 连接Redis
		connectRedis()
	}
	if LogrusConfigInstance.Enabled {
		// 设置log
		setLoggerInstance()
	}
	if ElasticConfigInstance.Enabled {
		// 连接elasticSearch
		connectElastic()
	}
}

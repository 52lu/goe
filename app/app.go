package app

import (
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"goe/app/common"
	"goe/app/config"
	. "goe/app/controllers"
	"net/http"
)

type App struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	Env  string `ini:"env"`
}

// --- 定义全局变量
var (
	RouteListInstance   = &RouteList{Route: map[string]interface{}{}}
	MysqlConfigInstance = &config.MysqlConfig{}
	BaseModel        = &common.BaseModel{}
	BusErrorInstance = &common.BusError{}
)

const (
	ConfigPath = "./app/config/" //配置文件目录
	EnvDev     = "dev"           //
	EnvLocal   = "local"
	EnvProd    = "prod"
	EnvPrepub  = "prod"
)

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
	// 注册路由
	app.registeredRoute()
	// 数据库连接
	app.connectMysql()
	// 启动服务
	fmt.Printf("Goe 启动成功！ Host:%s Port:%s \n", app.Host, app.Port)
	err := http.ListenAndServe(app.Host+":"+app.Port, RouteListInstance)
	if err != nil {
		fmt.Println("启动失败: " + err.Error())
	}
}

/**
 * @description: 加载
 * @user: Mr.LiuQH
 * @receiver app App
 * @date 2021-02-03 20:54:45
 */
func (app *App) registeredRoute() {
	// 注册路由
	RouteListInstance.AddRoute("user", &UserController{})
}

/**
 * @description: 加载配置文件
 * @user: Mr.LiuQH
 * @receiver app
 * @date 2021-02-04 15:29:17
 */
func (app *App) loadConfig() {
	iniPath := ConfigPath + app.Env + ".ini"
	fmt.Println("加载配置文件: " + iniPath)
	cfg, err := ini.Load(iniPath)
	BusErrorInstance.ThrowError(err)
	err = cfg.Section("app").MapTo(app)
	// 加载mysql配置
	err = cfg.Section("mysql").MapTo(MysqlConfigInstance)
}

/**
 * @description: 设置数据库连接
 * @user: Mr.LiuQH
 * @receiver app
 * @date 2021-02-04 17:06:29
 */
func (app *App) connectMysql() {
	// 连接数据库
	// 用户名:密码@tcp(IP:port)/数据库?charset=utf8
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		MysqlConfigInstance.UserName, MysqlConfigInstance.Password, MysqlConfigInstance.Host, MysqlConfigInstance.Port,
		MysqlConfigInstance.Database, MysqlConfigInstance.Charset)
	db, err := sql.Open("mysql", dataSourceName)
	BusErrorInstance.ThrowError(err)
	BaseModel.DB = db
}

package app

import (
	"fmt"
	"github.com/go-ini/ini"
	"goe/app/common"
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
	AppRouteInstance = &AppRoute{Route: map[string]interface{}{}}
	MysqlConfigInstance = &common.MysqlConfig{}
)

const (
	ConfigPath = "./app/config/" //配置文件目录
	EnvDev    = "dev" //
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
	//AppRouteInstance = &AppRoute{Route: map[string]interface{}{}}
	// 加载配置文件
	app.loadConfig()
	// 注册路由
	app.registeredRoute()
	// 启动服务
	fmt.Printf("Goe 启动成功！ Host:%s Port:%s \n", app.Host, app.Port)
	err := http.ListenAndServe(app.Host+":"+app.Port, AppRouteInstance)
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
	AppRouteInstance.AddRoute("user", &UserController{})
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
	if err != nil {
		panic(err.Error())
	}
	err = cfg.Section("app").MapTo(app)
	// 加载mysql配置
	err = cfg.Section("mysql").MapTo(MysqlConfigInstance)

}

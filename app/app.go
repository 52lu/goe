package app

import (
	"fmt"
	. "goe/app/controllers"
	"net/http"
)

type App struct {
	Host string
	Port string
}

// --- 定义全局变量
var (
	AppRouteInstance *AppRoute
)

/**
 * @description: 启动服务
 * @user: Mr.LiuQH
 * @receiver app App
 * @date 2021-02-03 10:21:20
 */
func (app *App) Start() {
	AppRouteInstance = &AppRoute{Route: map[string]interface{}{}}
	// 注册路由
	app.RegisteredRoute()
	// 启动服务
	err := http.ListenAndServe(app.Host+":"+app.Port, AppRouteInstance)
	if err != nil {
		fmt.Println("启动失败: " + err.Error())
	}
	fmt.Printf("Goe 启动成功！ Host:%s Port:%s \n",app.Host,app.Port)
}
/**
 * @description: 加载
 * @user: Mr.LiuQH
 * @receiver app App
 * @date 2021-02-03 20:54:45
 */
func (app *App) RegisteredRoute() {
	// 注册路由
	AppRouteInstance.AddRoute("user", &UserController{})
}

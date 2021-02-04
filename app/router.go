package app

import (
	"net/http"
	"reflect"
	"strings"
)


// 定义路由储存组
type AppRoute struct {
	Route map[string]interface{}
}


/**
 * @description: 注册路由
 * @user: Mr.LiuQH
 * @receiver receiver RouteConfig
 * @date 2021-02-03 11:48:03
 */
func (receiver *AppRoute) AddRoute(pattern string, controller interface{}) {
	receiver.Route[pattern] = controller
}

/**
 * @description: 处理接收请求
 * @user: Mr.LiuQH
 * @receiver receiver
 * @param w
 * @param r
 * @date 2021-02-03 15:35:26
 */
func (receiver *AppRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 捕获请求过程中的错误
	defer BusErrorInstance.CatchError()
	// 路由转发
	routeForWard(w, r)
	return
}

/**
 * @description: 路由转发
 * @user: Mr.LiuQH
 * @param urlPath
 * @return string
 * @return string
 * @date 2021-02-03 15:29:09
 */
func routeForWard(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if urlPath == "/favicon.ico" {
		return
	}
	// 路由解析
	split := strings.Split(r.URL.Path, "/")
	controller, methodName := split[1], strings.Title(split[2])
	if controller == "" || methodName == "" {
		http.NotFound(w, r)
		return
	}
	// 匹配路由
	controllerStruct, ok := AppRouteInstance.Route[controller]
	if !ok {
		http.NotFound(w, r)
		return
	}
	controllerValType := reflect.ValueOf(controllerStruct)
	// 保存请求信息到控制器基类
	controllerValType.Elem().FieldByName("Response").Set(reflect.ValueOf(w))
	controllerValType.Elem().FieldByName("Request").Set(reflect.ValueOf(r))
	// 保存到业务错误类里面
	BusErrorInstance.Response = w
	// 判断方法是否存在
	valid := controllerValType.MethodByName(methodName).IsValid()
	if !valid {
		http.NotFound(w, r)
		return
	}
	// 调用方法
	controllerValType.MethodByName(methodName).Call(nil)
}

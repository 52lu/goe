package common

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)



// 定义路由储存组
type RouteList struct {
	//Route map[string]interface{}
	Route map[string]map[string]interface{}
}


/**
 * @description: 注册路由
 * @user: Mr.LiuQH
 * @receiver receiver RouteConfig
 * @date 2021-02-03 11:48:03
 */
func (receiver *RouteList) AddRoute(version,pattern string, controller interface{}) {
	if  receiver.Route[version] == nil {
		receiver.Route[version] = make(map[string]interface{})
	}
	receiver.Route[version][pattern] = controller
}

/**
 * @description: 处理接收请求
 * @user: Mr.LiuQH
 * @receiver receiver
 * @param w
 * @param r
 * @date 2021-02-03 15:35:26
 */
func (receiver *RouteList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	// 解析参数
	parseError := r.ParseForm()
	if parseError != nil {
		panic("参数解析失败:" + parseError.Error())
	}
	// 获取版本号
	version := getVersion(r)
	fmt.Println("version:" + version)
	//  匹配路由
	controllerValType := matchControllerObj(version,controller,methodName)
	// 保存请求上下文到控制器基类
	controllerValType.Elem().FieldByName("Response").Set(reflect.ValueOf(w))
	controllerValType.Elem().FieldByName("Request").Set(reflect.ValueOf(r))
	BusErrorInstance.Response = w
	// 调用方法
	controllerValType.MethodByName(methodName).Call(nil)
}


/**
 * @description: 获取版本信息
 * @user: Mr.LiuQH
 * @param r
 * @return string
 * @date 2021-02-19 17:56:16
 */
func getVersion( r *http.Request) string  {
	var version string
	if r.Method == "GET" {
		version = r.FormValue("ver")
	} else if r.Method == "POST" {
		version = r.PostFormValue("ver")
	}
	if version == "" {
		version = r.Header.Get("ver")
	}
	if version == "" {
		panic(ReqParamVersionLost)
	}
	return version
}
/**
 * @description: 匹配路由
 * @user: Mr.LiuQH
 * @param version
 * @param controller
 * @param methodName
 * @return interface{}
 * @date 2021-02-19 18:35:07
 */
func matchControllerObj(version,controller,methodName string) reflect.Value  {
	vGroup,ok := RouteListInstance.Route[version]
	if !ok {
		panic(ReqVersionNotExist)
	}
	verNumStr := strings.Trim(version,"ver")
	verNum, _ := strconv.Atoi(verNumStr)
	// 匹配路由
	controllerStruct, ok := vGroup[controller]
	if !ok && verNum > 1 {
		newVer := "v"+ strconv.Itoa(verNum-1)
		return matchControllerObj(newVer,controller,methodName)
	}
	controllerValType := reflect.ValueOf(controllerStruct)
	// 判断方法是否存在
	valid := controllerValType.MethodByName(methodName).IsValid()
	if !valid {
		panic(ReqMethodNotFoundMsg)

	}
	return controllerValType
}
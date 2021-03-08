### 1.背景介绍
使用Go开发Web API框架，验证平时学习成果，在实践中深入学习。框架支持多版本API,集成Gorm、配置文件解析、后续随着学习深入，继续集成其他服务。
#### 1.1 集成功能

| 集成服务 | 笔记摘要                                                     |
| -------- | ------------------------------------------------------------ |
| Gorm     | [框架开发(二):集成Gorm及使用](http://liuqh.icu/2021/02/02/go/advanced/2-goe-gorm/) |
| Go-Redis | [框架开发(三):集成Go-Redis及使用](http://liuqh.icu/2021/02/03/go/advanced/3-goe-redis/) |
|go-ini    |       无                                                      |
|logrus    |  [框架开发(四):集成Logrus及使用](http://liuqh.icu/2021/02/02/go/advanced/2-goe-gorm/)                                                            |





### 2.目录介绍

```properties
├── README.md
├── app
│   ├── app.go //主程序文件
│   ├── common // 公共文件目录
│   │   ├── base_controller.go // 控制器基类
│   │   ├── base_model.go // 模型基类
│   │   ├── error.go // 错误处理
│   │   ├── global.go // 全局变量
│   │   ├── cli.go // 输出到控制台样式
│   │   ├── response_code.go // api返回code定义
│   │   └── router.go // 路由解析
│   ├── config // 配置目录
│   │   ├── dev.ini // 测试环境配置(默认加载)
│   │   ├── local.ini // 本地环境配置
│   │   └── mysql_config.go
│   ├── controllers // 控制器目录
│   │   ├── v1 
│   │   └── v2
│   ├── models // 数据库模型文件
│   │   └── user.go
│   └── static // 静态资源目录
│       └── a.png
├── go.mod
├── go.sum
├── main.go // 启动文件
└── test // 测试目录
```

### 3. 启动流程

![image-20210222151006966](https://gitee.com/QingHui/picGo-img-bed/raw/master/img/image-20210222151006966.png)



### 4.路由设计

框架设计的路由支持`Api`多版本,以及版本间的兼容性。

#### 4.1 期望效果

![image-20210222153003507](https://gitee.com/QingHui/picGo-img-bed/raw/master/img/image-20210222153003507.png)



#### 4.2 定义路由结构体

```go
// 定义路由储存组
type RouteList struct {
	//Route map[版本][控制器]处理器
	Route map[string]map[string]interface{}
}
```

#### 4.3 添加注册路由方法

```go
/**
 * @description: 注册路由
 * @user: Mr.LiuQH
 * @receiver receiver RouteConfig
 */
func (receiver *RouteList) AddRoute(version, pattern string, controller interface{}) {
	if receiver.Route[version] == nil {
		receiver.Route[version] = make(map[string]interface{})
	}
	receiver.Route[version][pattern] = controller
}
```

#### 4.4 路由注册

在`app.go`主程序文件中，通过`import _ 包路径`的方式，执行控制器包中的`init()`,从而达到自动注册路由。

**文件:**  `app/app.go`

```go
package app
import (
	...
	_ "goe/app/controllers/v1"
	_ "goe/app/controllers/v2"
  ...
)
```

**文件:**  `app/controller/v1/Test.go`

```go
/**
 * @Author Mr.LiuQH
 * @Description V1版本，Test控制器
 * @Date 2021/2/19 4:21 下午
 **/
package v1
import "goe/app/common"
type TestController struct {
	common.BaseController
}
func init() {
  // 注册路由
	common.RouteListInstance.AddRoute("v1","test",&TestController{})
}
func (t TestController) Hello() error {
	return t.Error("v1 hello")
}
func (t TestController) Run() error {
	return t.Error("v1 Run")
}
```

**文件:**  `app/controller/v2/Test.go`

```go
/**
 * @Author Mr.LiuQH
 * @Description V2版本，Test控制器
 * @Date 2021/2/19 4:21 下午
 **/
package v2
import "goe/app/common"
type TestController struct {
	common.BaseController
}
func init() {
	common.RouteListInstance.AddRoute("v2","test",&TestController{})
}
func (t TestController) Hello() error  {
	return t.Error("v2 hello")
}
```

#### 4.4 路由解析

每个`Http`请求，最终都会被`ServeHTTP`处理。

**文件:**  `app/common/router.go`

```go
/**
 * @description: 处理接收请求
 * @user: Mr.LiuQH
 * @receiver receiver
 * @param w
 * @param r
 */
func (receiver *RouteList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 初始化业务类属性
	BusErrorInstance.Response = w
	// 捕获请求过程中的错误
	defer BusErrorInstance.CatchError()
	// 静态路由解析
	isStaticReq := staticForWard(w, r)
	if !isStaticReq {
		// 动态路由解析
		routeForWard(w, r)
	}
}
```

##### 1.静态路由解析

```go
/**
 * @description: 静态路由解析
 * @user: Mr.LiuQH
 * @param w
 * @param r
 * @return bool
 * @date 2021-02-22 11:36:37
 */
func staticForWard(w http.ResponseWriter, r *http.Request) bool {
	if strings.HasPrefix(r.URL.String(), "/static/") {
		prefixHttp := http.StripPrefix("/static/", http.FileServer(http.Dir("app/static")))
		prefixHttp.ServeHTTP(w, r)
		return true
	}
	return false
}
```

##### 2.动态路由解析

根据控制器和版本，找到对应的处理器，然后通过反射，调用对应的方法.

```go
/**
 * @description: 动态路由转发
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
	controllerValType := matchControllerObj(version, controller, methodName)
	// 保存请求上下文到控制器基类
	controllerValType.Elem().FieldByName("Response").Set(reflect.ValueOf(w))
	controllerValType.Elem().FieldByName("Request").Set(reflect.ValueOf(r))
	BusErrorInstance.Response = w
	// 调用方法
	controllerValType.MethodByName(methodName).Call(nil)
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
func matchControllerObj(version, controller, methodName string) reflect.Value {
	fmt.Printf("进入匹配路由: version:%s controller:%s  methodName:%s \n", version, controller, methodName)
	vGroup, ok := RouteListInstance.Route[version]
	if !ok {
		panic(ReqVersionNotExist)
	}
	verNumStr := strings.Trim(version, "ver")
	verNum, _ := strconv.Atoi(verNumStr)
	// 匹配路由
	controllerStruct, ok := vGroup[controller]
  // 当前版本没有，则找下个版本
	if !ok && verNum > 1 {
		newVer := "v" + strconv.Itoa(verNum-1)
		return matchControllerObj(newVer, controller, methodName)
	}
	controllerValType := reflect.ValueOf(controllerStruct)
	// 判断方法是否存在
	valid := controllerValType.MethodByName(methodName).IsValid()
	if !valid {
		panic(ReqMethodNotFoundMsg)
	}
	return controllerValType
}
```


### 5.启动
```bigquery
➜  go run main.go 
+--------+----------------------+
|  KEY   |        VALUE         |
+--------+----------------------+
| Config | ./app/config/dev.ini |
| Host   | 127.0.0.1            |
| Port   | 9923                 |
+--------+----------------------+
```

### 6.开发笔记
- [框架开发(一):目录介绍和路由设计](http://liuqh.icu/2021/02/01/go/advanced/1-goe/)
- [框架开发(二):集成Gorm及使用](http://liuqh.icu/2021/02/02/go/advanced/2-goe-gorm/)
- [框架开发(三):集成Go-Redis及使用](http://liuqh.icu/2021/02/03/go/advanced/3-goe-redis/)


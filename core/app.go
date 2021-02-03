package core

import (
	"fmt"
	"net/http"
)

type AppServer struct {
	Host string
	Port string
}
/**
 * @description: 启动服务
 * @user: Mr.LiuQH
 * @receiver app AppServer
 * @date 2021-02-03 10:21:20
 */
func (app AppServer)Start()  {
	routeConfig := &RouteConfig{}
	err := http.ListenAndServe(app.Host+":"+app.Port, routeConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
}




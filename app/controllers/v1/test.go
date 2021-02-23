/**
 * @Author Mr.LiuQH
 * @Description V1版本，Test控制器
 * @Date 2021/2/19 4:21 下午
 **/
package v1

import (
	"goe/app/common"
)

type TestController struct {
	common.BaseController
}
func init() {
	common.RouteListInstance.AddRoute("v1","test",&TestController{})
}
func (t TestController) Hello() error {
	return t.Error("v1 hello")
}
func (t TestController) Run() error {
	//ctx := context.Background()
	//common.RedisClient.Set()
	return t.Error("v1 Run")
}
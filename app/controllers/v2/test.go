/**
 * @Author Mr.LiuQH
 * @Description V2版本，Test控制器
 * @Date 2021/2/19 4:21 下午
 **/
package v2

import (
	"goe/app/common"
)

type TestController struct {
	common.BaseController
}

func init() {
	common.RouteListInstance.AddRoute("v2","test",&TestController{})
}
func (t TestController) Hello() error  {
	return t.Error("v2 hello")
}


/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/22 9:43 上午
 **/
package v1

import "goe/app/common"

type UserController struct {
	common.BaseController
}
func init() {
	common.RouteListInstance.AddRoute("v1","user",&UserController{})
}

func (receiver UserController) List() error {
	return receiver.Error("数据为空!")
}
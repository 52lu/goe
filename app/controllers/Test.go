/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/19 4:29 下午
 **/
package controllers

import "goe/app/common"

func init() {
	common.RouteListInstance.AddRoute("test",&TestController{})
}
type TestController struct {
	common.BaseController
}

func (c TestController) Echo() error {
	return c.Error("Hello")
}
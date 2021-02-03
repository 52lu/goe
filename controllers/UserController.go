/**
 * @Author Mr.LiuQH
 * @Description 用户相关接口
 * @Date 2021/1/28 9:40 下午
 **/
package controllers

import "goe/core"

type UserController struct {
	core.BaseController
}

func (user *UserController) login() map[string]string {
	m := map[string]string{
		"name":"张飒",
		"home":"北京",
	}
	return m
}

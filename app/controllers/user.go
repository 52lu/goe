/**
 * @Author Mr.LiuQH
 * @Description 用户相关接口
 * @Date 2021/1/28 9:40 下午
 **/
package controllers

import (
	"goe/app/common"
)

type UserController struct {
	common.BaseController
}

type LoginReturn struct {
	Name string
	Home string
}

func (user *UserController) Login() error {
	m := LoginReturn{
		"李四",
		"北京",
	}
	return user.Success(m)
}

func (user *UserController) Config() error {

	return user.Success(nil)
}
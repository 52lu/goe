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
/**
 * @description: 账号密码登录
 * @user: Mr.LiuQH
 * @receiver user
 * @return error
 * @date 2021-02-04 18:23:23
 */
func (user *UserController) Login() error {
	m := LoginReturn{
		"李四",
		"北京",
	}
	return user.Success(m)
}



/**
 * @description: 注册
 * @user: Mr.LiuQH
 * @receiver user
 * @return error
 * @date 2021-02-04 18:23:48
 */
func (user *UserController) Register() error {
	panic("抛错测试")
	//return user.Success(nil)
	return nil
}
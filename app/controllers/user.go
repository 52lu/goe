/**
 * @Author Mr.LiuQH
 * @Description 用户相关接口
 * @Date 2021/1/28 9:40 下午
 **/
package controllers

import (
	"crypto/md5"
	"fmt"
	"goe/app/common"
	"goe/app/models"
	"strconv"
)

//func init() {
//	common.RouteListInstance.AddRoute("user", &UserController{})
//}

type UserController struct {
	common.BaseController
}

type LoginReturn struct {
	UserName string
	Password string
}
/**
 * @description: 账号密码登录
 * @user: Mr.LiuQH
 * @receiver user
 * @return error
 * @date 2021-02-04 18:23:23
 */
func (uc *UserController) Login() error {
	userName := uc.GetParam("userName")
	password := uc.GetParam("password")
	if userName == "" || password == "" {
		 return uc.Error("参数不能为空!")
	}
	m := LoginReturn{
		userName,
		password,
	}
	return uc.Success(m)
}
/**
 * @description: 查询用户
 * @user: Mr.LiuQH
 * @receiver uc
 * @return error
 * @date 2021-02-19 15:27:03
 */
func (uc *UserController) GetUser() error  {
	id := uc.GetParam("id")
	if id == "" {
		return uc.Error("id不能为空")
	}
	userOne := &models.User{}
	idNumber, _ := strconv.Atoi(id)
	userOne.FindById(idNumber)
	return uc.Success(userOne)
}


/**
 * @description: 注册
 * @user: Mr.LiuQH
 * @receiver uc
 * @return error
 * @date 2021-02-04 18:23:48
 */
func (uc *UserController) Register() error {
	nickName := uc.GetParam("nickName")
	email := uc.GetParam("email")
	mobile := uc.GetParam("mobile")
	birthday := uc.GetParam("birthday")
	notEmptyParam := []string{nickName,email,mobile,birthday}
	for _,v := range notEmptyParam {
		if v == "" {
			return uc.Error(fmt.Sprintf("%s不能为空!",v))
		}
	}
	// 判断用户是否存在
	userExist := &models.User{}
	userExist.FindByMobile(mobile)
	if userExist.ID != 0 {
		return uc.Error(fmt.Sprintf("手机号%s已经存在!",mobile))
	}
	// 插入新用户
	userOne := &models.User{
		NickName: nickName,
		Email: email,
		Mobile: mobile,
		Birthday: birthday,
		Status: 1,
		Password: fmt.Sprintf("%x",md5.Sum([]byte(mobile))),
	}
	userOne.Add()
	// 入库
	return uc.Success(userOne)
}
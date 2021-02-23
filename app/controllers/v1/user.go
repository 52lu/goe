/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/22 9:43 上午
 **/
package v1

import (
	"crypto/md5"
	"fmt"
	"goe/app/common"
	"goe/app/models"
	"strconv"
	"time"
)

type UserController struct {
	common.BaseController
}

func init() {
	common.RouteListInstance.AddRoute("v1", "user", &UserController{})
}

/**
 * @description: 查询
 * @user: Mr.LiuQH
 * @receiver u UserController
 * @return error
 * @date 2021-02-23 09:57:29
 */
func (uc UserController) GetUser() error {
	uid := uc.GetParam("uid")
	phone := uc.GetParam("phone")
	userModel := &models.User{}
	if uid != "" {
		id, _ := strconv.Atoi(uid)
		userModel.FindById(id)
	} else if phone != "" {
		userModel.FindByMobile(phone)
	} else {
		return uc.Error("查询条件不能为空!")
	}
	return uc.Success(userModel)
}

/**
 * @description: 注册
 * @user: Mr.LiuQH
 * @receiver uc
 * @return error
 * @date 2021-02-04 18:23:48
 */
func (uc UserController) Register() error {
	nickName := uc.GetParam("nickName")
	email := uc.GetParam("email")
	mobile := uc.GetParam("mobile")
	birthday := uc.GetParam("birthday")
	notEmptyParam := []string{nickName, email, mobile, birthday}
	for _, v := range notEmptyParam {
		if v == "" {
			return uc.Error(fmt.Sprintf("%s不能为空!", v))
		}
	}
	// 判断用户是否存在
	userExist := &models.User{}
	userExist.FindByMobile(mobile)
	if userExist.ID != 0 {
		return uc.Error(fmt.Sprintf("手机号%s已经存在!", mobile))
	}

	location, _ := time.LoadLocation("Asia/Shanghai")
	birthdayTime, _ := time.ParseInLocation("2006-01-02", birthday, location)

	// 插入新用户
	userOne := &models.User{
		NickName: nickName,
		Email:    email,
		Mobile:   mobile,
		Birthday: common.DateTime(birthdayTime),
		Status:   1,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(mobile))),
	}
	userOne.Add()
	// 入库
	return uc.Success(userOne)
}

/**
 * @description: 更新用户信息
 * @user: Mr.LiuQH
 * @receiver uc UserController
 * @return error
 * @date 2021-02-23 10:08:36
 */
func (uc UserController) Update() error {
	uid := uc.GetParam("uid")
	name := uc.GetParam("name")
	phone := uc.GetParam("phone")
	userModel := &models.User{}
	id, _ := strconv.Atoi(uid)
	userModel.FindById(id)
	userUpdate := models.User{
		NickName: name,
		Mobile:   phone,
	}
	userModel.UpdateStatus(userUpdate)
	return uc.Success(userModel)
}

/**
 * @Author Mr.LiuQH
 * @Description 用户相关的模型
 * @Date 2021/2/4 5:56 下午
 **/
package models

import (
	"goe/app/common"
)

type User struct {
	common.BaseModel
	NickName string          `json:"nickName"`
	Password string          `json:"password"`
	Email    string          `json:"email"`
	Mobile   string          `json:"mobile"`
	Gender   int8            `json:"gender"`
	Birthday common.DateTime `json:"birthday"`
	Status   int8            `json:"status"`
}

// 根据主键查询
func (um *User) FindById(id int) {
	if result := common.GormDBClient.First(um, id); result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}

// 添加单条记录
func (um *User) Add() {
	if result := common.GormDBClient.Create(um); result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}

// 根据条件查询
func (um *User) FindByMobile(mobile string) {
	if result := common.GormDBClient.Where("mobile=?", mobile).Find(um); result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}

/**
 * @description: 更新用户信息
 * @user: Mr.LiuQH
 * @receiver um
 * @param user
 * @date 2021-02-23 10:03:14
 */
func (um *User) UpdateStatus(user User) {
	if result := common.GormDBClient.Model(um).Updates(user); result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}
/**
 * @description: 删除用户(软删除)
 * @user: Mr.LiuQH
 * @receiver um
 * @date 2021-02-23 15:27:43
 */
func (um *User) DelUser(id int)  {
	if result := common.GormDBClient.Delete(um,id);result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}
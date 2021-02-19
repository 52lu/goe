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
	NickName string
	Password string
	Email    string
	Mobile   string
	Gender   int8
	Birthday string
	Status   int8
}

// 根据主键查询
func (um *User) FindById(id int)  {
	if result := common.GormDBClient.First(um,id); result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}
// 创建
func (um *User) Add()  {
	if result := common.GormDBClient.Create(um);result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}
// 根据条件查询
func (um *User) FindByMobile( mobile string)  {
	if result := common.GormDBClient.Where("mobile=?",mobile).Find(um);result.Error != nil {
		common.BusErrorInstance.ThrowError(result.Error)
	}
}




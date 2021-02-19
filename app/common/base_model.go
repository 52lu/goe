/**
 * @Author Mr.LiuQH
 * @Description 模型父类
 * @Date 2021/2/8 10:49 上午
 **/
package common

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedAt int
	UpdatedAt int
}

/**
 * @Author Mr.LiuQH
 * @Description 模型父类
 * @Date 2021/2/8 10:49 上午
 **/
package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey;autoincrement;not null" json:"id"`
	CreatedAt int `json:"createdAt"`
	UpdatedAt int `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) error  {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = DateTime(t1)
	return err
}

func (t DateTime) MarshalJSON() ([]byte,error)  {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}
func (t DateTime) Value() (driver.Value, error) {
	// DateTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}
func (t *DateTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = DateTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}
func (t *DateTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
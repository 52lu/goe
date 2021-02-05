/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/4 5:57 下午
 **/
package common

import "database/sql"

type BaseModel struct {
	DB *sql.DB
}

func (b *BaseModel) findOne(column []string, )  {
	//b.DB.Prepare()
}
/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/4 5:59 下午
 **/
package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BusError struct {
	Response    http.ResponseWriter
}

/**
 * @description: 捕获错误
 * @user: Mr.LiuQH
 * @receiver b
 * @date 2021-02-04 18:00:11
 */
func (b *BusError) CatchError() {
	// 捕获全局错误
	err := recover()
	msg := fmt.Sprintf("运行异常: %v", err)
	if err != nil {
		if b.Response != nil {
			marshal, _ := json.Marshal(map[string]interface{}{
				"code": 500,
				"msg":  msg,
			})
			b.Response.Write(marshal)
			return
		}
		fmt.Println(msg)
	}
}

/**
 * @description: 抛出错误
 * @user: Mr.LiuQH
 * @receiver b
 * @param err
 * @date 2021-02-04 18:01:05
 */
func (b *BusError) ThrowError(err error)  {
	if err != nil {
		panic("处理失败:"+err.Error())
	}
}
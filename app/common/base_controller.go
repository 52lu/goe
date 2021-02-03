/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/3 10:26 上午
 **/
package common

import (
	"encoding/json"
	"net/http"
)

type BaseController struct {
	Response    http.ResponseWriter
	Request     *http.Request
}
type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/**
 * @description: 处理成功
 * @user: Mr.LiuQH
 * @receiver b BaseController
 * @param data
 * @date 2021-02-03 18:45:09
 */
func (b BaseController)Success(data interface{}) error  {
	result := ApiResponse{
		Code: ReqSuccess,
		Msg:  ReqSuccessMsg,
		Data: data,
	}
	marshal, _ := json.Marshal(result)
	b.Response.Write([]byte(marshal))
	return nil
}

/**
 * @description: 处理失败
 * @user: Mr.LiuQH
 * @receiver b BaseController
 * @param msg
 * @date 2021-02-03 18:45:20
 */
func (b BaseController)Error(msg string) error  {
	result := ApiResponse{
		Code: ReqError,
		Msg:  msg,
	}
	marshal, _ := json.Marshal(result)
	b.Response.Write([]byte(marshal))
	return nil
}
/**
 * @Author Mr.LiuQH
 * @Description 控制器基类
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
type apiResponse struct {
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
	result := apiResponse{
		Code: ReqSuccess,
		Msg:  ReqSuccessMsg,
		Data: data,
	}
	marshal, _ := json.Marshal(result)
	b.Response.Write(marshal)
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
	result := apiResponse{
		Code: ReqError,
		Msg:  msg,
	}
	marshal, _ := json.Marshal(result)
	b.Response.Write(marshal)
	return nil
}
/**
 * @description: 获取GET参数
 * @user: Mr.LiuQH
 * @receiver b BaseController
 * @param key
 * @return string
 * @date 2021-02-05 16:13:54
 */
func (b BaseController)GetParam(key string)string  {
	return b.Request.FormValue(key)
}
/**
 * @description: 获取Post参数
 * @user: Mr.LiuQH
 * @receiver b BaseController
 * @param key
 * @return string
 * @date 2021-02-05 16:13:54
 */
func (b BaseController)PostParam(key string)string  {
	return b.Request.PostFormValue(key)
}


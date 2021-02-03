/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/3 10:26 上午
 **/
package core

import "net/http"

type BaseController struct {
	response http.ResponseWriter
	request  *http.Request
}

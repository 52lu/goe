/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/3/8 2:15 下午
 **/
package v1

import (
	"github.com/sirupsen/logrus"
	. "goe/app/common"
)

type LogController struct {
	BaseController
}

func init() {
	RouteListInstance.AddRoute("v1","log",&LogController{})

}
/**
 * @description: 测试logrus使用
 * @user: Mr.LiuQH
 * @receiver l LogController
 * @return error
 * @date 2021-03-08 14:22:45
 */
func (l LogController) Test() error  {
	// 简短消息记录
	LoggerClient.Trace("这是Trace,日志信息")
	LoggerClient.Debug("这是Debug,日志信息")
	LoggerClient.Info("这是Info,日志信息")
	LoggerClient.Error("这是Error,日志信息")
	LoggerClient.Warn("这是Warn,日志信息")
	//记录结构化数据
	LoggerClient.WithFields(logrus.Fields{
		"uerName":"zhangsan",
		"age":28,
		"pice":500.12,
		"likes":[]string{"游戏","旅游"},
	}).Info("记录结构化数据")
	return  l.Success(nil)
}
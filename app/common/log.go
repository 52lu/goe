/**
 * @Author Mr.LiuQH
 * @Description 日志相关
 * @Date 2021/3/8 2:35 下午
 **/
package common

import (
	"fmt"
	goFileRotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"strings"
	"time"
)

type CustomizeFormat struct {
}

/**
 * @description: 自定义格式输出
 * @user: Mr.LiuQH
 * @receiver c CustomizeFormat
 * @param entry
 * @return []byte
 * @return error
 * @date 2021-03-08 15:53:29
 */
func (c CustomizeFormat) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] %s \n",
		time.Now().Local().Format("2006-01-02 15:04:05"),
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)
	return []byte(msg), nil
}

/**
 * @description: 写入日志并分割
 * @user: Mr.LiuQH
 * @param save
 * @date 2021-03-08 15:16:32
 */
func Log2FileByClass() {
	lfhook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: splitConfig("debug"),
		logrus.InfoLevel:  splitConfig("info"),
		logrus.WarnLevel:  splitConfig("warn"),
		logrus.ErrorLevel: splitConfig("error"),
		logrus.FatalLevel: splitConfig("fatal"),
		logrus.PanicLevel: splitConfig("painc"),
	}, LoggerClient.Formatter)
	LoggerClient.AddHook(lfhook)
}

/**
 * @description: 分割文件配置
 * @user: Mr.LiuQH
 * @param level
 * @return *goFileRotatelogs.RotateLogs
 * @date 2021-03-08 15:49:08
 */
func splitConfig(level string) *goFileRotatelogs.RotateLogs {
	// 拼凑日志名
	logFile := path.Join(LogrusConfigInstance.Path, level)
	logs, err := goFileRotatelogs.New(
		// 文件名
		logFile+"-"+LogrusConfigInstance.Suffix,
		// 生成软链，指向最新日志文件
		goFileRotatelogs.WithLinkName(logFile),
	)
	BusErrorInstance.ThrowError(err)
	return logs
}

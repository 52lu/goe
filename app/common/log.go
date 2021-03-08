/**
 * @Author Mr.LiuQH
 * @Description 日志相关
 * @Date 2021/3/8 2:35 下午
 **/
package common

import (
	goFileRotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)
/**
 * @description: TODO
 * @user: Mr.LiuQH
 * @param save
 * @date 2021-03-08 15:16:32
 */
func Log2FileByClass() {
	lfhook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: write("debug"),
		logrus.InfoLevel:  write("info"),
		logrus.WarnLevel:  write("warn"),
		logrus.ErrorLevel: write("error"),
		logrus.FatalLevel: write("fatal"),
		logrus.PanicLevel: write("painc"),
	}, LoggerClient.Formatter)
	LoggerClient.AddHook(lfhook)
}

func write(level string) *goFileRotatelogs.RotateLogs {
	// 拼凑日志名
	logFile := path.Join(LogrusConfigInstance.Path, level)
	//解析日志切割间隔时间
	splitTime, err := time.ParseDuration(LogrusConfigInstance.SplitTime)
	logs, err := goFileRotatelogs.New(
		// 文件名
		logFile+"-"+LogrusConfigInstance.Suffix,
		// 生成软链，指向最新日志文件
		goFileRotatelogs.WithLinkName(logFile),
		//文件最大保存份数
		goFileRotatelogs.WithRotationCount(int(LogrusConfigInstance.ClassSaveNum)),
		// 文件最大保存时间
		goFileRotatelogs.WithMaxAge(time.Minute * 3),
		//日志切割时间间隔
		goFileRotatelogs.WithRotationTime(splitTime),
	)
	BusErrorInstance.ThrowError(err)
	return logs

}

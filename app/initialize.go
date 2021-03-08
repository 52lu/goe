/**
 * @Author Mr.LiuQH
 * @Description 初始化文件
 * @Date 2021/3/8 11:25 上午
 **/
package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	. "goe/app/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

/**
 * @description: 连接mysql(todo 需要设置成单例)
 * @user: Mr.LiuQH
 * @date 2021-02-08 10:21:42
 */
func connectMysql() {
	// 用户名:密码@tcp(IP:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		MysqlConfigInstance.UserName, MysqlConfigInstance.Password, MysqlConfigInstance.Host, MysqlConfigInstance.Port,
		MysqlConfigInstance.Database, MysqlConfigInstance.Charset, MysqlConfigInstance.ParseTime, MysqlConfigInstance.Loc)

	// 连接额外配置信息
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   MysqlConfigInstance.TablePre, //表前缀
			SingularTable: true,                         //使用单数表名，启用该选项时，`User` 的表名应该是 `user`而不是users
		},
	}
	// 打印SQL设置
	if MysqlConfigInstance.PrintSqlLog {
		slowSqlTime, err := time.ParseDuration(MysqlConfigInstance.SlowSqlTime)
		BusErrorInstance.ThrowError(err)
		loggerNew := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: slowSqlTime, //慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true, // 彩色打印开启
		})
		gormConfig.Logger = loggerNew
	}
	// 建立连接
	GormDBClient, err = gorm.Open(mysql.Open(dataSourceName), gormConfig)
	BusErrorInstance.ThrowError(err)
	// 设置连接池信息
	db, err2 := GormDBClient.DB()
	BusErrorInstance.ThrowError(err2)
	// 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(MysqlConfigInstance.MaxIdleConn)
	// 设置打开数据库连接的最大数量
	db.SetMaxOpenConns(MysqlConfigInstance.MaxOpenConn)
	// 设置了连接可复用的最大时间
	duration, err := time.ParseDuration(MysqlConfigInstance.MaxLifeTime)
	BusErrorInstance.ThrowError(err)
	db.SetConnMaxLifetime(duration)
	// 打印SQL配置信息
	//marshal, _ := json.Marshal(db.Stats())
	//fmt.Printf("数据库配置: %s \n", marshal)
}

/**
 * @description: 连接Redis
 * @user: Mr.LiuQH
 * @date 2021-02-23 17:06:27
 */
func connectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisConfigInstance.Host + ":" + RedisConfigInstance.Port,
		Password: RedisConfigInstance.Password,  // 密码
		DB:       RedisConfigInstance.DefaultDB, // 默认数据库
		PoolSize: RedisConfigInstance.PoolSize,  // 连接池大小
	})
	// 设置连接超时时间
	duration, err := time.ParseDuration(RedisConfigInstance.TimeOut)
	ctx, cancelFunc := context.WithTimeout(context.Background(), duration)
	defer cancelFunc()
	_, err = RedisClient.Ping(ctx).Result()
	BusErrorInstance.ThrowError(err)
}

/**
 * @description: 设置全局日志
 * @user: Mr.LiuQH
 * @date 2021-03-08 14:11:07
 */
func setLoggerInstance() {
	// 设置日志级别
	var level logrus.Level
	err := level.UnmarshalText([]byte(LogrusConfigInstance.Level))
	BusErrorInstance.ThrowError(err)
	LoggerClient.SetLevel(level)
	// 设置日志格式
	if LogrusConfigInstance.Formatter == "json" {
		LoggerClient.SetFormatter(&logrus.JSONFormatter{})
	} else if LogrusConfigInstance.Formatter == "text" {
		LoggerClient.SetFormatter(&logrus.TextFormatter{})
	} else if LogrusConfigInstance.Formatter == "customize" {
		LoggerClient.SetFormatter(&CustomizeFormat{})
	} else {
		BusErrorInstance.ThrowError(errors.New("log formatter must json|text|customize"))
	}
	// 打开日志记录的行数；true:开启，false:关闭。默认关闭
	if LogrusConfigInstance.ReportCaller {
		LoggerClient.SetReportCaller(LogrusConfigInstance.ReportCaller)
	}
	// 设置日志输出方式
	switch LogrusConfigInstance.OutputType {
	case "1":
		// 控制台
		LoggerClient.SetOutput(os.Stdout)
	case "2":
		// 文件
		Log2FileByClass()
	default:
		// 默认写到控制台
		LoggerClient.SetOutput(os.Stdout)
	}
}

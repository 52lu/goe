/**
 * @Author Mr.LiuQH
 * @Description 全局变量
 * @Date 2021/2/8 10:11 上午
 **/
package common

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"goe/app/config"
	"gorm.io/gorm"
)

// 全局业务变量单例
var (
	RouteListInstance = &RouteList{Route: make(map[string]map[string]interface{})}
	BusErrorInstance  = &BusError{}
	CliInfoInstance   = CliInfo{}
)

// 全局客户端变量
var (
	GormDBClient *gorm.DB
	RedisClient  *redis.Client
	LoggerClient = logrus.New()
)

// 全局配置变量
var (
	MysqlConfigInstance  = &config.MysqlConfig{}
	RedisConfigInstance  = &config.RedisConfig{}
	LogrusConfigInstance = &config.LogrusConfig{}
)

// 全局常量
const (
	ConfigPath = "./app/config/" //配置文件目录
	StaticPath = "./app/static/" //静态资源文件目录
	EnvDev     = "dev"           //
	EnvLocal   = "local"
	EnvProd    = "prod"
	EnvPrepub  = "prod"
)

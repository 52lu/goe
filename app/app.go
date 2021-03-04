package app

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	. "goe/app/common"
	_ "goe/app/controllers/v1"
	_ "goe/app/controllers/v2"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	Env  string `ini:"env"`
}

var err error

/**
 * @description: 启动服务
 * @user: Mr.LiuQH
 * @receiver app App
 * @date 2021-02-03 10:21:20
 */
func (app *App) Start() {
	// 捕获启动时错误
	defer BusErrorInstance.CatchError()
	// 加载配置文件
	app.loadConfig()
	// 初始化数据库连接
	app.initializeDB()
	// 设置打印的信息
	CliInfoInstance.Host = app.Host
	CliInfoInstance.Port = app.Port
	CliInfoInstance.PrintRunMsg()
	// 启动服务
	err := http.ListenAndServe(app.Host+":"+app.Port, RouteListInstance)
	if err != nil {
		fmt.Println("启动失败: " + err.Error())
	}
}

/**
 * @description: 加载配置文件
 * @user: Mr.LiuQH
 * @receiver app
 * @date 2021-02-04 15:29:17
 */
func (app *App) loadConfig() {
	iniPath := ConfigPath + app.Env + ".ini"
	CliInfoInstance.ConfigFile = iniPath
	cfg, err := ini.Load(iniPath)
	BusErrorInstance.ThrowError(err)
	err = cfg.Section("app").MapTo(app)
	// 加载mysql配置
	err = cfg.Section("mysql").MapTo(MysqlConfigInstance)
	// 加载redis配置
	err = cfg.Section("redis").MapTo(RedisConfigInstance)
	if err != nil {
		BusErrorInstance.ThrowError(err)
	}
}

/**
 * @description: 设置数据库连接
 * @user: Mr.LiuQH
 * @receiver app
 * @date 2021-02-04 17:06:29
 */
func (app *App) initializeDB() {
	// 连接Mysql
	connectMysql()
	// 连接Redis
	connectRedis()
}

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
		Password: RedisConfigInstance.Password, // 密码
		DB:       RedisConfigInstance.DefaultDB, // 默认数据库
		PoolSize: RedisConfigInstance.PoolSize, // 连接池大小
	})
	// 设置连接超时时间
	duration, err := time.ParseDuration(RedisConfigInstance.TimeOut)
	ctx, cancelFunc := context.WithTimeout(context.Background(), duration)
	defer cancelFunc()
	result, err := RedisClient.Ping(ctx).Result()
	fmt.Println("redis: " + result)
	BusErrorInstance.ThrowError(err)
}

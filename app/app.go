package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-ini/ini"
	. "goe/app/common"
	_ "goe/app/controllers"
	_ "goe/app/controllers/v1"
	_ "goe/app/controllers/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
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
	// 注册路由
	//app.registeredRoute()
	// 初始化数据库连接
	app.initializeDB()
	// 启动服务
	fmt.Printf("Goe 启动成功！ Host:%s Port:%s \n", app.Host, app.Port)
	fmt.Printf("路由映射: %+v\n", RouteListInstance.Route)
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
	fmt.Println("加载配置文件: " + iniPath)
	cfg, err := ini.Load(iniPath)
	BusErrorInstance.ThrowError(err)
	err = cfg.Section("app").MapTo(app)
	// 加载mysql配置
	err = cfg.Section("mysql").MapTo(MysqlConfigInstance)
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
}

/**
 * @description: 连接mysql(todo 需要设置成单例)
 * @user: Mr.LiuQH
 * @date 2021-02-08 10:21:42
 */
func connectMysql()  {
	// 用户名:密码@tcp(IP:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		MysqlConfigInstance.UserName, MysqlConfigInstance.Password, MysqlConfigInstance.Host, MysqlConfigInstance.Port,
		MysqlConfigInstance.Database, MysqlConfigInstance.Charset, MysqlConfigInstance.ParseTime, MysqlConfigInstance.Loc)
	GormDBClient, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: MysqlConfigInstance.TablePre, //表前缀
			SingularTable: true, //使用单数表名，启用该选项时，`User` 的表名应该是 `user`而不是users
		},
	})
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
	marshal, _ := json.Marshal(db.Stats())
	fmt.Printf("数据库配置: %s \n" , marshal)
}
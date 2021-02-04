/**
 * @Author Mr.LiuQH
 * @Description 数据库相关的配置
 * @Date 2021/2/4 4:14 下午
 **/
package common

/**
 * @description: mysql配置
 * @user: Mr.LiuQH
 */
type MysqlConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Database string `ini:"database"`
	UserName string `ini:"userName"`
	Password string `ini:"password"`
}

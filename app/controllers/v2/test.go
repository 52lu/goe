/**
 * @Author Mr.LiuQH
 * @Description V2版本，Test控制器
 * @Date 2021/2/19 4:21 下午
 **/
package v2

import (
	"github.com/go-redis/redis/v8"
	"goe/app/common"
	"golang.org/x/net/context"
	"time"
)

type TestController struct {
	common.BaseController
}

func init() {
	common.RouteListInstance.AddRoute("v2","test",&TestController{})
}
func (t TestController) Hello() error  {
	return t.Error("v2 hello")
}
/**
 * @description: redis设置
 * @user: Mr.LiuQH
 * @receiver t TestController
 * @return error
 * @date 2021-02-23 18:06:48
 */
func (t TestController) ReSet() error  {
	key := t.GetParam("key")
	val := t.GetParam("val")
	ctx := context.Background()
	err := common.RedisClient.Set(ctx, key, val, time.Second * 60).Err()
	if err != nil {
		return t.Error("Set: " + err.Error())
	}
	return t.Success(nil)
}
/**
 * @description: redis获取
 * @user: Mr.LiuQH
 * @receiver t TestController
 * @return error
 * @date 2021-02-23 18:06:58
 */
func (t TestController) ReGet() error  {
	key := t.GetParam("key")
	ctx := context.Background()
	result, err := common.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return t.Error(key + " not exist")
	} else if err != nil {
		return t.Error("Get: " + err.Error())
	}
	return t.Success(result)
}
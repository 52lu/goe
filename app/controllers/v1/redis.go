/**
 * @Author Mr.LiuQH
 * @Description 测试redis相关的操作
 * @Date 2021/2/23 6:19 下午
 **/
package v1

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	. "goe/app/common"
	"time"
)

type RedisController struct {
	BaseController
}

func init() {
	RouteListInstance.AddRoute("v1","redis",&RedisController{})
}
var ctx = context.Background()
/**
 * @description: redis设置
 * @user: Mr.LiuQH
 * @receiver t TestController
 * @return error
 * @date 2021-02-23 18:06:48
 */
func (r RedisController) Normal() error  {
	opType := r.GetParam("opType")
	key := r.GetParam("key")
	if opType == "1" {
		// Set
		val := r.GetParam("val")
		err := RedisClient.Set(ctx, key, val, time.Second * 60).Err()
		if err != nil {
			return r.Error("Set: " + err.Error())
		}
		return r.Success(nil)
	} else  {
		// Get
		result, err := RedisClient.Get(ctx, key).Result()
		if err == redis.Nil {
			return r.Error(key + " not exist")
		} else if err != nil {
			return r.Error("Get: " + err.Error())
		}
		return r.Success(result)
	}
}


/**
* @description: 有序集合添加
* @user: Mr.LiuQH
* @receiver r RedisController
* @date 2021-02-23 18:22:41
*/
func (r RedisController) SortSet() error  {
	opType := r.GetParam("opType")
	key := r.GetParam("key")
	if opType == "1" {
		// 有序集合添加
		zs := []*redis.Z{
			{Member: "小张",Score: 88},
			{Member: "小李",Score: 90},
			{Member: "小明",Score: 80},
			{Member: "小英",Score: 70},
			{Member: "小赵",Score: 95},
			{Member: "小王",Score: 75},
			{Member: "笨蛋",Score: 40},
		}
		result, err := RedisClient.ZAdd(ctx, key, zs...).Result()
		if err != nil {
			return r.Error(err.Error())
		}
		return r.Success(result)
	} else  {
		resultMap := make(map[string]interface{})
		// 获取成员数
		val := RedisClient.ZCard(ctx, key).Val()
		resultMap["获取成员数"] = val
		// 获取指定分数区间的成员数
		resultMap["70分-90分成员数"] = RedisClient.ZCount(ctx, key, "70", "90").Val()
		// 返回有序集中指定分数区间内的成员，分数从高到低排序
		result, _ := RedisClient.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
			Min: "0", Max: "100", Offset: 0, Count: 3,
		}).Result()
		resultMap["分数前三排名"] = result
		// 返回分数值
		f, _ := RedisClient.ZScore(ctx, key, "小张").Result()
		resultMap["小张的分数"] = f
		// 给笨蛋加60分
		f2, err := RedisClient.ZIncrBy(ctx, key, 60.0, "笨蛋").Result()
		fmt.Println(err)
		fmt.Printf("新分数:%v \n",f2)
		result2, _ := RedisClient.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
			Min: "0", Max: "100", Offset: 0, Count: 3,
		}).Result()
		resultMap["调整分后排名"] = result2
		return r.Success(resultMap)
	}
}

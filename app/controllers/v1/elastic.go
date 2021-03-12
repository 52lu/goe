/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/3/10 6:47 下午
 **/
package v1

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	. "goe/app/common"
	"math/rand"
	"strconv"
	"time"
)

type ElasticController struct {
	BaseController
}

func init() {
	RouteListInstance.AddRoute("v1", "elastic", &ElasticController{})
}

// 定义用户mapping
const userMapping = `
{
    "mappings":{
        "properties":{
            "name":{
                "type":"keyword"
            },
            "age":{
                "type":"byte"
            },
            "phone":{
                "type":"text"
            },
            "birth":{
                "type":"date"
            },
            "height":{
                "type":"float"
            },
            "smoke":{
                "type":"boolean"
            },
            "home":{
                "type":"geo_point"
            }
        }
    }
}`

// 定义索引名
const userIndex = "go-user"

var esCtx = context.Background()

// 创建索引
func (e ElasticController) CreateIndex() error {
	exist, err := ElasticClient.IndexExists(userIndex).Do(esCtx)
	if err != nil {
		return e.Error(err.Error())
	}
	if exist {
		return e.Error("索引已经存在，无需重复创建!")
	}
	// 创建索引
	res, err := ElasticClient.CreateIndex(userIndex).BodyString(userMapping).Do(esCtx)
	LoggerClient.WithFields(logrus.Fields{
		"res":res,
	}).Info("es返回")
	if err != nil {
		return e.Error(err.Error())
	}
	return e.Success(res)
}

type User struct {
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	Phone  string    `json:"phone"`
	Birth  time.Time `json:"birth"`
	Height float32   `json:"height"`
	Smoke  bool      `json:"smoke"`
	Home   string    `json:"home"`
}

// 添加文档
func (e ElasticController) AddOneDoc() error {
	// 定义上下文
	loc, err := time.LoadLocation("Local")
	birth, err := time.ParseInLocation("2006-01-02", "1991-04-25", loc)
	userOne := User{"张三", 23, "17600000000", birth, 170.5, false,"41.40338,2.17403"}
	do, err := ElasticClient.Index().Index(userIndex).Id(strconv.FormatInt(time.Now().UnixNano(),10)).BodyJson(userOne).Do(esCtx)
	if err != nil {
		return e.Error(err.Error())
	}
	return e.Success(do)
}
// 批量添加
func (e ElasticController) BatchAddDoc() error {
	userBulk := ElasticClient.Bulk().Index(userIndex)
	loc, _ := time.LoadLocation("Local")
	// 生日
	birthSlice := []string{"1991-04-25","1990-01-15","1989-11-05","1988-01-25","1994-10-12"}
	// 姓名
	nameSlice := []string{"李四","张飞","赵云","关羽","刘备"}
	rand.Seed(time.Now().Unix())
	for i := 1; i < 20; i++ {
		birth, _ := time.ParseInLocation("2006-01-02", birthSlice[rand.Intn(len(birthSlice))], loc)
		height, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float32()+175.0), 32)
		user := User{
			nameSlice[rand.Intn(len(nameSlice))],
			rand.Intn(10) + 18,
			"1760000000" +strconv.Itoa(i),
			birth,
			float32(height),
			false,
			"41.40338,2.17403",
		}
		fmt.Println(user, userBulk)
		doc := elastic.NewBulkIndexRequest().Id(strconv.FormatInt(time.Now().UnixNano(), 10)).Doc(user)
		userBulk.Add(doc)
	}
	if userBulk.NumberOfActions() < 0 {
		return e.Error("没有要保存的数据")
	}
	if _, err := userBulk.Do(esCtx);err != nil{
		return e.Error(err.Error())
	}
	return e.Success(nil)
}


// 查询
func (e ElasticController) Get() error {
	//ElasticClient.Get().Index(userIndex).Id().
	//ElasticClient.Get().Index(userIndex).
	return nil
}

func (e ElasticController) Del() error {
	return nil
}
func (e ElasticController) Update() error {
	return nil
}

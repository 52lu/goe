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
	"reflect"
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
                "type":"text"
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
		"res": res,
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
	userOne := User{"张三", 23, "17600000000", birth, 170.5, false, "41.40338,2.17403"}
	do, err := ElasticClient.Index().Index(userIndex).Id(strconv.FormatInt(time.Now().UnixNano(), 10)).BodyJson(userOne).Do(esCtx)
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
	birthSlice := []string{"1991-04-25", "1990-01-15", "1989-11-05", "1988-01-25", "1994-10-12"}
	// 姓名
	nameSlice := []string{"李四", "张飞", "赵云", "关羽", "刘备"}
	rand.Seed(time.Now().Unix())
	for i := 1; i < 20; i++ {
		birth, _ := time.ParseInLocation("2006-01-02", birthSlice[rand.Intn(len(birthSlice))], loc)
		height, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float32()+175.0), 32)
		user := User{
			nameSlice[rand.Intn(len(nameSlice))],
			rand.Intn(10) + 18,
			"1760000000" + strconv.Itoa(i),
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
	if _, err := userBulk.Do(esCtx); err != nil {
		return e.Error(err.Error())
	}
	return e.Success(nil)
}

// 根据名字查询
func (e ElasticController) Get() error {
	name := e.GetParam("name")
	list, err := ElasticClient.Search().Index(userIndex).
		Query(elastic.NewMatchQuery("name", name)).
		Sort("age", true). //根据age字段，正序
		From(0). // 从第几条开始取
		Size(20). // 每页数量
		Pretty(true). //返回json格式
		Do(context.Background())
	if err != nil {
		return e.Error(err.Error())
	}
	var userDto User
	var userList []User
	for _, v := range list.Each(reflect.TypeOf(userDto)) {
		tmp := v.(User)
		userList = append(userList, tmp)
	}
	// 遍历获取结果
	return e.Success(userList)
}

// 删除
func (e ElasticController) Del() error {
	id := e.GetParam("id")
	name := e.GetParam("name")
	age := e.GetParam("age")
	var do interface{}
	var err error
	if id != "" {
		// 根据主键删除
		do, err = ElasticClient.Delete().Index(userIndex).Id(id).Do(context.Background())
	} else if name != "" && age != "" {
		// 根据名称和年龄删除
		do, err = ElasticClient.DeleteByQuery().Index(userIndex).
			Query(elastic.NewMatchQuery("name", name)).
			Query(elastic.NewTermQuery("age", age)).
			Do(context.Background())
	} else {
		return e.Error("缺少参数信息")
	}
	if err != nil {
		return e.Error(err.Error())
	}
	return e.Success(do)
}

// 更新
func (e ElasticController) Update() error {
	id := e.GetParam("id")
	phone := e.GetParam("phone")
	var do interface{}
	var err error
	if id != "" {
		// 根据主键更新
		do, err = ElasticClient.Update().Index(userIndex).Id(id).
			Script(elastic.NewScript("ctx._source.age=73")).
			Script(elastic.NewScript("ctx._source.phone='110110110110'")).
			Do(context.Background())
	} else if phone != "" {
		// 根据非ID条件更新
		do, err = ElasticClient.UpdateByQuery(userIndex).Query(elastic.NewTermQuery("phone", phone)).
			// 通过脚本更新字段name
			Script(elastic.NewScript("ctx._source.name='龙少爷'")).
			// 如果文档版本冲突继续执行
			ProceedOnVersionConflict().
			Do(context.Background())
	} else {
		return e.Error("缺少参数信息")
	}
	if err != nil {
		return e.Error(err.Error())
	}
	return e.Success(do)
}

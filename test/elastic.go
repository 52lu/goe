/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/3 3:03 下午
 **/
package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
	"time"
)

func TestElastic() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://114.116.248.223:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHeaders(http.Header{
			"X-Caller-Id": []string{"..."},
		}),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	do, err := client.DeleteByQuery().Index("go-user").
		Query(elastic.NewMatchQuery("name", "关羽")).
		Query(elastic.NewTermQuery("age", 22)).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(do.Took)
}

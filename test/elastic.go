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

func main2() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
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
	res, err := client.ClusterHealth().Index("test").Level("shards").Pretty(true).Do(context.TODO())

	fmt.Println(res)
}

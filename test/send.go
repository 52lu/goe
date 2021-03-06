package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)
var waitGroup sync.WaitGroup
var sendTimes = 30
//电信诈骗网址
var url string = "http://p.qiao.baidu.com/cps/chat"
func tt() {
	times := 10
	waitGroup.Add(times)
	for i := 0; i < times; i++ {
		go Req(i)
	}
	waitGroup.Wait()
}

func Req(g int)  {
	for i := 0; i < sendTimes; i++ {
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err.Error())
		}
		if response == nil {
			continue
		}
		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("协程%d:%d  StatusCode: %v Length: %v \n",g,i,response.StatusCode,len(r))
		response.Body.Close()
	}
	waitGroup.Done()
}

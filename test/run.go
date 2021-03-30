/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/3/30 11:16 上午
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	TestElastic()
}




func testAfter()  {
	//ch := make(chan string)
	//go func(ch chan string) {
	//	ch <- time.Now().String()
	//}(ch)
	//
	//select {
	//case str := <-ch:
	//	fmt.Printf("接收的内容: %v\n",str)
	//case time.After():
	//
	//}
}

func testTicker()  {
	// 创建定时器，间隔设置每秒
	ticker := time.NewTicker(time.Second)
	// 启动一个协程，打印定时器里面的时间
	go func(ticker *time.Ticker) {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ticker.C)
		}
		// 关闭定时器
		ticker.Stop()
	}(ticker)
	// 手动阻塞
	time.Sleep(3 * time.Second)
	fmt.Println("end")
}

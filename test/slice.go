/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/3/12 11:27 上午
 **/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//str := []string{"a","b","v","s","i","o","p"}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println("Unix=>",strconv.FormatInt(time.Now().Unix(),10))
		fmt.Println("UnixNano=>",strconv.FormatInt(time.Now().UnixNano(),10))
	}


}

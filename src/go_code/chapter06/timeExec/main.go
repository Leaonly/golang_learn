package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0;i<100000;i++{
		str +=  "hello" + strconv.Itoa(i)
	}
}

func main() {
	//统计时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Println("执行时间为: ",end-start,"秒")
	
}
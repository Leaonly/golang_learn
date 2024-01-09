package main

import (
	"fmt"
	"time"
)

func main() {
	// 日期相关函数
	now := time.Now()
	fmt.Println("now=",now)
	//通过now获取年月日等
	fmt.Println("now年=",now.Year())

	

}
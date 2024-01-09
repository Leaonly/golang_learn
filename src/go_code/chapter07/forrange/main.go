package main

import (
	"fmt"
)

func test02(arr *[3]int){
	(*arr)[1] = 88
}

func main(){
	// 遍历数组
	heroes := [...]string{"超人","蝙蝠侠","闪电侠"}

	for i, v := range heroes{
		fmt.Printf("i=%v, v=%v\n",i,v)
	}

	// 可以只取值
	for _, v := range heroes{
		fmt.Printf("v=%v\n",v)
	}

	// 数组也是值传递，可以用指针参数传递修改值
	arr := [3]int{11,22,33}
	fmt.Printf("arr地址=%p\n",&arr)
	test02(&arr)
	fmt.Println("arr=",arr)
}
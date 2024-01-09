package main

import (
	"fmt"
)

func main(){
	// 方式1 引用数组定义切片
	var arr [5]int = [...]int{1,2,3,4,5}
	var slice1 = arr[1:3]
	fmt.Println("arr=",arr)
	fmt.Println("slice1=",slice1)

	// 方式2.演示切片的使用 make
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	//默认值为0
	fmt.Println(slice2)

	// 方式3 切片直接定义数组
	var strSlice []string = []string{"tom","jack","marry"}
	fmt.Println(strSlice)
	fmt.Println(len(strSlice))
	fmt.Println(cap(strSlice))
}
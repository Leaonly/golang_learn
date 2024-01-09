package main

import (
	"fmt"
)

func main(){
	// 常规for遍历切片
	var arr[5]int = [...]int{10,20,30,40,50}
	slice := arr[1:4]
	for i := 0 ; i<len(slice);i++{
		fmt.Printf("slice[%v]=%v \n",i,slice[i])
	}

	// 使用for range
	for i,v := range slice{
		fmt.Printf("range i=%v v=%v\n",i,v)
	}

	// 切片用append追加元素
	var slice3 []int = []int{100,200,300}
	slice3 = append(slice3, 400,600)
	fmt.Println("slice3=",slice3)
	// 追加切片
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3=",slice3)
}
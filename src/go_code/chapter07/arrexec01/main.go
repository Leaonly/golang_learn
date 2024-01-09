package main

import (
	"fmt"
)

func main(){
	/*
	// 将A-Z字母放到数组并打印
	arrBytes := [26]byte{}
	for i:=0;i<=25;i++ {
		arrBytes[i] = 'A' + byte(i) 	//需要将i转成byte

	}
	for _,v := range arrBytes {
		fmt.Println(string(v))
	}
	*/

	// 判断数组中最大值
	intArr := [...]int{1,-1,20,90,600,202}
	maxValue := 0
	for _,v := range intArr{
		if v > maxValue {
			maxValue = v
		}
	}
	fmt.Println("最大值为: ",maxValue)
}
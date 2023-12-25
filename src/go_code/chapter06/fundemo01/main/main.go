package main

import (
	"fmt"
	"go_code/chapter06/fundemo01/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'
	// 调用utils函数
	result := utils.Cal(n1,n2,operator)
	fmt.Println("result=",result)

	//调用utils变量
	fmt.Println("调用函数Num1为：",utils.Num1)
}
package main

import (
	"fmt"
	"math/rand"
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
	intArr := [...]int{1,-1,20,90,600,202,900}
	maxValue := 0
	minValue := intArr[0]
	sum := 0
	for _,v := range intArr{
		if v > maxValue {
			maxValue = v
		}
		if v < minValue {
			minValue = v
		}
		sum += v
	}
	fmt.Println("最大值为: ",maxValue)
	fmt.Println("最小值为: ",minValue)
	fmt.Println("平均值为: ",float64(sum)/float64(len(intArr)))

	// 生成5个随机数并翻转打印
	intArr3 := [5]int{}
	for i:=0 ; i<len(intArr3);i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println("随机数为: ",intArr3)

	// for i := len(intArr3)-1;i>=0;i--{
	// 	fmt.Println("翻转打印为: ",intArr3[i])
	// }
	// initVal := 0
	// lenArr := len(intArr3) - 1
	// for i,v := range intArr3{
	// 	initVal = intArr3[lenArr] - i
	// 	intArr3[lenArr - i]= v
	// 	intArr3[i] = initVal
	// }
	// fmt.Println("翻转打印为: ",intArr3)

	tmp := 0
	for i:=0;i < len(intArr3)/2;i++{
		tmp = intArr3[len(intArr3)- 1 - i]
		intArr3[len(intArr3)- 1 - i] = intArr3[i]
		intArr3[i] = tmp
	}
	fmt.Println("翻转打印为: ",intArr3)

	

}
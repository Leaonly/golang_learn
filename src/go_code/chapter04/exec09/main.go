package main

import (
	"fmt"
)

func main() {
	// var n1 float64 = 12.1
	// var n2 float64 = 22.2
	// if n1 > 10.0 && n2 > 20.0 {
	// 	fmt.Println(n1 + n2)
	// }

	// var n3 int32 = 215
	// var n4 int32 = 100
	// var n5 int32 = n3 + n4
	// if n5%3 == 0 && n5%5 == 0 {
	// 	fmt.Println("能被整除ok")
	// }

	var year int32
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)

	if year % 4 == 0 &&  year % 100 != 0 || year % 400 == 0 {
		fmt.Printf("%v是闰年!",year)
	}
}

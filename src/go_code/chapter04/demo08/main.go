package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	if age > 18{
		fmt.Println("你已成年!")
	} else {
		fmt.Println("未成年禁止吸烟！")
	}
}
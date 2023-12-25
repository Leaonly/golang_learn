package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Println("请输入名字：")
	fmt.Scanln(&name)
	fmt.Println("年龄：")
	fmt.Scanln(&age)
	fmt.Println("薪水：")
	fmt.Scanln(&sal)
	fmt.Println("是否通过：")
	fmt.Scanln(&isPass)

	fmt.Printf(" 名字：%v \n 年龄：%v\n 薪水：%v\n 是否通过：%v", name, age, sal, isPass)

	// 方法2 交互全部参数在一行
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

}

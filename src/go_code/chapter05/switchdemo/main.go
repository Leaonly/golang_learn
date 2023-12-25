package main

import "fmt"

func main() {
	// 定义一个字符，switch分支执行

	// var aa string
	// fmt.Println("请输入一个字母: ")
	// fmt.Scanln(&aa)

	// switch aa {
	// case "a":
	// 	fmt.Println("aaaaaaaaaa")
	// case "b":
	// 	fmt.Println("bbbb")
	// default:
	// 	fmt.Println("xxxxxxxxxxxxx")
	// }

	// switch后也可以不带表达式
	var score int = 100
	switch  {
	case score > 90:
		fmt.Println("very good")
	case score < 90 && score >= 60:
		fmt.Println("good")
	case score < 60:
		fmt.Println("bad")
	}

	// switch穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
	}



}

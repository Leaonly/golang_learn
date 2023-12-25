package main

import (
	"fmt"
)

func test() bool{
	fmt.Println("test...")
	return true
}

func main(){

	var i int = 10

	if i > 9 && test(){
		fmt.Println("ok...")
	}

	// 前面条件已为true，所以只打印hello
	if i > 9 || test() {
		fmt.Println("hello...")
	}

	// var i int = 97
	// var b int = i / 7
	// var c int = 97 - b * 7
	// fmt.Printf("97天假为：%v星期，%v天\n",b,c)

	// var huat float32 = 150.0
	// var shh float32 = 5.0/9.0*(huat-100)
	// fmt.Println("摄氏度为：", shh)




}



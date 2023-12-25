package main 


import (
	"fmt"
)

func main (){
	var a int = 10
	var b int = 20
	// 交换a, b变量值方法1
	// a, b = b, a
	// fmt.Println(a, b)

	// 方法2
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

	fmt.Println("地址：",&a)
	var ptr *int = &a
	fmt.Println("指针值：",*ptr)
}

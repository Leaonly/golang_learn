package main

import "fmt"

func main() {
	/*
	// 打印99乘法表
	var result int = 9
	for n1 := 1; n1 <= result; n1++ {
		for n2 := 1; n2 <= n1; n2++ {
			fmt.Printf("%v * %v = %v \t", n1, n2, n1*n2)
		}
		fmt.Println()
	}

	// 选题1公式（身高-108）*2 =体重，10左右浮动
	var height int 
	var weight int
	fmt.Println("请输入身高(cm)：")
	fmt.Scanln(&height)
	fmt.Println("请输入体重(kg)：")
	fmt.Scanln(&weight)
	
	if (height - 108) * 2 >= weight * 2 - 10 && (height - 108) * 2 <= weight *2 + 10 {
		fmt.Println("体重达标")
	} else {
		fmt.Println("体重不达标")
	}
	*/

	//选题2
	var score int
	fmt.Println("请输入分数：")
	fmt.Scanln(&score)

	switch  {
	case score > 90:
		fmt.Println("优秀")
	case score <= 90 && score > 60:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}
}





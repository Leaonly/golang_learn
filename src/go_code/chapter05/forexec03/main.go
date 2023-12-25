package main

import "fmt"

func main() {

	// 打印矩形*
	for i := 1;i <= 5;i++ {
		for j := 1;j <= i ;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 打印三角形*
	var total int = 9
	for i := 1;i <= total;i++ {
		// 打印空格
		for k := 1;k <=total - i;k++ {
			fmt.Print(" ")
		}
		for j := 1;j <= i *2 -1 ;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
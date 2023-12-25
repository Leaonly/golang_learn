package main

import (
	"fmt"
)
// 用递归函数表示斐波那契数， 第N个数的 前一位-1 前两位-2的和
func fbn(n int) int {
	if (n == 1 || n == 2) {
		return 1
	} else {
		return fbn(n -1) + fbn(n -2)
	}
}

func main (){
	fmt.Println("res=",fbn(8))
}
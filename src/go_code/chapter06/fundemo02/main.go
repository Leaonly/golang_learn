package main

import (
	"fmt"
)

func getSumAndSub(n1 int, n2 int) (int,int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	//调用函数，返回多个值
	res1, res2 := getSumAndSub(20, 10)
	fmt.Printf("返回为res1=%v ,res2=%v",res1,res2)
}

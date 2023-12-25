package main

import (
	"fmt"
)

// 用递归函数表示斐波那契数， 第N个数的 前一位-1 前两位-2的和
func chiTaoZi(day int) int {
	if day == 10 {
		return 1
	} else {
		return (chiTaoZi(day+1) + 1) * 2
	}
}

func test02(n1 int) int {
	n1 = n1 + 10
	return n1
}

// 函数内修改变量，用指针地址传递修改
func test03(n2 *int) {
	*n2 = *n2 + 10
}

func main() {
	fmt.Printf("第N天桃子为%v个\n", chiTaoZi(6))

	n1 := 20
	test02(n1)
	fmt.Print("====", test02(n1), "\n")

	n2 := 20
	test03(&n2)
	fmt.Println("函数指针传递n2为 ", n2)
}

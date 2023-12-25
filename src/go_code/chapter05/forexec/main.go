package main

import "fmt"

func main() {
	// 打印0-100所有9的倍数整数的个数与和
	a := 0
	b := 0
	for i := 0;i <= 100;i++ {
		if i % 9 == 0 {
			a++
			b += i
		}
	}
	fmt.Println("100里9的倍数个数为: ",a)
	fmt.Println("100里9的倍数之和为: ",b)

	//练习2
	var n int = 10
	for i := 0;i <= n;i++ {
		fmt.Printf("%v + %v = %v \n",i,n-i,n)
	}

	//练习3
	var j int = 1
	for {
		fmt.Printf("hello, %v",j)
		j++
		if j > 10{
			break
		}
	}
}
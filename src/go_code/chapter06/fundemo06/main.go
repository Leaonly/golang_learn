package main

import (
	"fmt"
)

var (
	//Fun1是一个全局匿名函数
	Fun1 = func (n1 int, n2 int)int{
		return n1 * n2
	}
)

func main(){
	//匿名函数
	//案例1 
	res1 := func (n1 int,n2 int)int{
		return n1 + n2 
	}(10, 20)
	fmt.Println("res1=",res1)

	//案例2
	a := func (n1 int,n2 int)int{
		return n1 - n2 
	}
	fmt.Println("res2=",a(10,20))

	//全局匿名函数调用
	res4 := Fun1(4, 9)
	fmt.Println("res4=",res4)

}
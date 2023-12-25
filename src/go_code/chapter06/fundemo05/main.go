package main

import (
	"fmt"
)

func getSum (n1 int,n2 int)int{
	return n1 + n2
}

// 函数可以作为形参调用
func myFun(funvar func(int,int)int, num1 int,num2 int) int {
	return funvar(num1, num2)
}

// 自定义函数类型
type myFunType func(int ,int) int

func myFun3(funvar myFunType, num1 int,num2 int) int {
	return funvar(num1, num2)
}


// 自定义函数返回值命名
func getSumAndSub(n1 int,n2 int)(sum int,sub int){
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 案例4  可变参数args...要放到参数最后
func sum (n1 int,args... int) int{
	sum := 0
	// 遍历args
	for i:=0;i<len(args);i++{
		sum += args[i]
	}
	return sum
}

func main(){
	a := getSum
	fmt.Printf("a的类型是%T, getSum类型是%T \n",a,getSum)

	// 函数也是数据类型，可以赋值给变量
	res := a(10,40)
	fmt.Println("res=",res,"")

	// 函数可以作为形参调用res
	res1 := myFun(getSum,50,60)
	fmt.Println("函数形参调用值例子为 ",res1,"")

	// 自定义函数类型结果
	res3 := myFun3(getSum,500,600)
	fmt.Println("自定义函数类型例子值为 ",res3,"")

	// 自定义函数返回名
	a1, b1 := getSumAndSub(1,2)
	fmt.Printf("a=%v, b=%v \n",a1,b1)

	res4 := sum(10,20,100,300)
	fmt.Println("res4=",res4,"")

}
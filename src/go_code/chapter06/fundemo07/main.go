package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 累加器，返回的函数引用了外部变量会视作闭包，变量只会初始化一次
func AddUpper() func (int) int {
	var n int = 10
	var str = "hello"
	return func(i int) int {
		n = n + i
		str += strconv.Itoa(i)
		fmt.Println("str=",str)
		return n
	}
}

// 练习1
/*
1)编写一个函数 makesuffix(suffix string) 可以接收一个文件后名(比如jpg)，并返回一个闭包2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg ，如果
3)要求使用闭包的方式完成
4)strings.HasSuffix ，该函数可以判断某个字符串是否有指定的后缀。
*/
func makeSuffix(suffix string)func (string)string{
	
	return func (name string)string  {
		//判断后缀是.jpg，如果否则加上后缀
		if ! strings.HasSuffix(name,suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	//练习1调用
	f2 := makeSuffix(".jpg")
	fmt.Println("后缀处理后=",f2("winner"))
	fmt.Println("后缀处理后=",f2("wow.jpg"))
	
}
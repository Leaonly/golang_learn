package main

import (
	"fmt"
)
//方法快速入门
type Person struct{
	Name string
}

//1给Person结构体添加speak方法,输出 xxx是一个好人
func (p Person) speak(){
	fmt.Println(p.Name,"是一个good man")
}

//2给Person结构体添加jisuan 方法可以计算从1+..+1000的结果
func (p Person) jisuan(){
	res := 0
	for i:=0;i<=1000;i++{
		res += i
	}
	fmt.Println("计算结果1是=",res)
}

//3给Person结构体jisuan2方法,该方法可以接收一个数n，i计算从 1+..+n 的结果
func (p Person) jisuan2(n int){
	res := 0
	for i:=0;i<=n;i++{
		res += i
	}
	fmt.Println("计算结果2是=",res)
}

//4)给Person结构体添加getSum方法,可以计算两个数的和，并返回结果。
func (p Person) getSum(n1 int,n2 int)int{
	return n1 + n2
}

func main(){
	var p Person
	p.Name = "tom"
	p.speak()

	p.jisuan()
	p.jisuan2(10)
	sum := p.getSum(10,20)
	fmt.Println("两数和=",sum)
}

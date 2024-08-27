package main

import (
	"fmt"
)

func main(){
	var num int
	//常量在定义时需要赋值，且之后不能修改
	const tax int = 0 
	//tax = 10 (error的)
	fmt.Println(num, tax)
	//批量定义常量，逐个+1
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a,b,c,d)
}
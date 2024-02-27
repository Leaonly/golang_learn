package main

import (
	"fmt"
)

// 定义cat 结构体
type Cat struct {
	Name string
	Age int
	Color string
}

type Person struct{
	Name string
	Age  int
}

func main(){

	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "write"
	fmt.Println(cat1)

	// 结构体实例的创建
	// 方法1
	var p1 Person
	p1.Name = "p1-name"
	p1.Age = 10
	fmt.Println(p1)

	// 方法2
	p2 := Person{"p2-name",20}
	fmt.Println(p2)

	// 方法3
	var p3 *Person = new(Person)
	(*p3).Name = "p3-name"
	p3.Name = "P3-name-plus"  //同等上面，简写
	(*p3).Age = 30
	p3.Age = p3.Age*30
	fmt.Println(*p3)

	// 方法4
	var person *Person = &Person{}
	fmt.Println(*person)

}
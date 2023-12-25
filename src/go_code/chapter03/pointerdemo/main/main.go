package main

import (
	"fmt"
	// "model1"
	"go_code/chapter03/pointerdemo/model1"
)

// 指针类型
func main(){
	// 基本数据内存布局
	var i int = 10
	fmt.Println("i的内存地址是", &i)

	// ptr是指针变量，类型是*int，ptr本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("ptr的地址是%v\n",&ptr)
	fmt.Printf("ptr指向的值是%v\n",*ptr)

	// 练习
	var num int = 9
	fmt.Printf("num adress=%v\n",num)

	var ptr2 *int
	ptr2 = &num
	*ptr2 = 10 //这里修改时，会更改num的值
	fmt.Println("num =",num)   

	var str_3 int = 100
	fmt.Print(str_3)
 
	fmt.Println("导入函数是",model1.HeroName)

}
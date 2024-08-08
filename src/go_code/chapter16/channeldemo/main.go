package main

import (
	"fmt"

)

func main(){
	var intChan chan int
	intChan = make(chan int,3)

	fmt.Printf("intChan值=%v intChan本身地址=%p \n",intChan, &intChan)

	//向管道写入数据, 但不能超过其容量
	intChan<- 10
	num := 211
	intChan<- num
	//看管道长度和cap容量
	fmt.Printf("channel len=%v cap=%v \n",len(intChan), cap(intChan))


	//5.从管道读取数据
	var num2 int 
	num2 = <-intChan
	fmt.Println("num2=",num2)
	fmt.Printf("channel len=%v cap=%v \n",len(intChan), cap(intChan))

	//6.再没有协程情况下，如果数据全部取出，则报错deadlock

}
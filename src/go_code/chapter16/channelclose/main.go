package main

import (
	"fmt"

)

func main(){
	intChan := make(chan int, 3)
	intChan<- 100
	intChan<- 200
	//关闭管道，则不能写入数据到channel
	close(intChan)
	//intChan<- 300
	fmt.Println("ok~")
	//关闭后，读取数据是可以的
	n1:= <-intChan
	fmt.Println("n1=",n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i :=0;i<100;i++{
		intChan2<- i*2 //放入100个数据
	}
	//遍历, 之前要关闭管道，否则会出现deadlock错误
	close(intChan2)
	for  v := range intChan2{
		fmt.Println("v=",v)
	}
	
}
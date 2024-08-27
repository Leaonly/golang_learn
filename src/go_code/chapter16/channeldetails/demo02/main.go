package main

import (
	"fmt"
	"time"
)

func main(){
	//使用selet解决管道阻塞问题

	//定义一个管道，10个数据int
	intChan := make(chan int,10)
	for i :=0; i<10;i++{
		intChan<-i
	}
	//定义一个管道，5个数据string
	stringChan:=make(chan string, 5)
	for i := 0; i<5 ;i++{
		stringChan <- "hello" + fmt.Sprintf("%d",i)
	}

	//实际开发中如不确定什么时候关闭管道，可以用select方式解决
	for{
		select {
			//这里intChan一直没关闭，不阻塞而会deadlock
			//会自动下一个case匹配
		case v:= <-intChan:	
			fmt.Printf("从intChan读取的数据%d\n",v) 
		case v:= <-stringChan:
			fmt.Printf("从stringChan读取数据%s\n",v)
		default:
			fmt.Printf("都取不到，可以自己加入逻辑\n")
			time.Sleep(time.Second)
			return
		}
	}
}
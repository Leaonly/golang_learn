package main

import (
	"fmt"
	"sync"
	"time"
)

//计算1-200各个数的阶乘，放入map中
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	lock sync.Mutex
)

func test(n int) {
	res:=1
	for i := 1;i<n;i++{
		res *= i
	}

	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main(){
	//启用协程写入map
	for i := 1;i<=20;i++{
		go test(i)
	} 
	
	// 等待协程执行完，设置大概时间5秒
	time.Sleep(time.Second*5)
	//遍历map
	lock.Lock()
	for i,v := range myMap{
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}
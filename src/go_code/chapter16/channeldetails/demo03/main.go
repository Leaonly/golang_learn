package main

import (
	"fmt"
	"time"
)

func sayHello(){
	for i:=0; i<4; i++{
		time.Sleep(time.Second)
		fmt.Println("hello, world")
	}
}

func test(){
	// 这里用defer+recover捕获test抛出的panic
	defer func ()  {
		if err:= recover();err != nil{
			fmt.Println("test() 发生错误",err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main(){
	go sayHello()
	go test()

	for i := 0; i<5;i++{
		fmt.Println("main() ok=",i)
		time.Sleep(time.Second)

	}
}
package main

import (
	"fmt"
)

func main() {
	// 创建一个整型通道
	ch := make(chan int)

	// 启动一个 Goroutine 发送数据
	go func() {
		ch <- 42
	}()

	// 从通道接收数据
	result := <-ch

	// 打印接收到的数据
	fmt.Println("Received:", result)
}

package main

import (
	"fmt"
)

func main() {
	// 打印1-100奇数 ，用continue
	for i:=1;i<=100;i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
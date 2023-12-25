package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个随机数，看为99要多少次
	count := 0
	for {
		v1 := rand.Intn(100)
		count++
		if v1 == 99{
			break
		}
	}
	fmt.Printf("用了 %v 次",count)

}
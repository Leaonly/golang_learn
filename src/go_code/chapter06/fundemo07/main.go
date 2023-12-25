package main

import (
	"fmt"
)

// 累加器
func AddUpper() func (int) int {
	var n int = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
}
package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("test1 n=",n)
}

func test2(n int) {
	if n >2 {
		n--
		test2(n)
	} else {
		fmt.Println("test2 n=",n)
	}
}

func main() {
	test(4)
	test2(4)
}
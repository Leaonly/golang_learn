package main

import (
	"fmt"
)

// 函数传入值互换，用指针交换
func swap(n1 *int, n2 *int)  {
	t := *n1
	*n1 = *n2
	*n2 = t
	//或者 *n1, *n2 = *n2, *n1
	
}
func main() {
	a := 10
	b := 20
	swap(&a,&b)
	fmt.Printf("a=%v, b=%v",a,b)
}

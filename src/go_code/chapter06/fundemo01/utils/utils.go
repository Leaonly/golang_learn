package utils

import "fmt"

var Num1 int = 100

func Cal(n1 float64, n2 float64, operater byte) float64 {
	var res float64
	switch operater {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误。。。")
		
	}
	return res 
}
package main

import (
	"fmt"
)

/*
定义小小计算器结构体(Calcuator)，实现加减乘除四个功能实现
形式1:分四个方法完成:实现形式2:用一个方法搞定
*/

type Calcuator struct{
	Num1 float64
	Num2 float64
}

func (run *Calcuator) wow(aa byte) float64{
	res := 0.0
	switch aa {
	case '+':
		res = float64(run.Num1) + float64(run.Num2)
	case '-':
		res = float64(run.Num1) - float64(run.Num2)
	case '*':
		res = float64(run.Num1) * float64(run.Num2)
	case '/':
		res = float64(run.Num1) / float64(run.Num2)
	default:
		fmt.Println("输入运算符错误...")
	}
	return float64(res)
}

func main(){
	var bb Calcuator
	bb.Num1 = 10.4
	bb.Num2 = 4.2

	res1 := bb.wow('*')
	fmt.Printf("结果为: %v",fmt.Sprintf("%.2f",res1))

}
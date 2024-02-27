package main

import (
	"fmt"
)
/*
方法和函数区别
1.调用方式不一样
函数的调用方式:
方法的调用方式:
函数名(实参列表)
变量,方法名(实参列表)
2.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
3.对于方法(如struct的方法)，接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
*/
type Box struct {
	len float64
	width float64
	height float64
}

func (box *Box) getVolumn() float64{
	return  box.len * box.height * box.width
}

func main(){
	var vol = Box{
		len: 10,
		width: 20,
		height: 3,
	}
	fmt.Println("体积为: ",vol.getVolumn())

}



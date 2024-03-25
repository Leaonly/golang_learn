package main
import (
	"fmt"
)

//定义一个接口
type Shape interface{
	Area() float64
	Perimeter() float64
}

//实现接口的具体类型
type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle)Area() float64 {
	return r.width * r.height
}

func (r Rectangle)Perimeter()float64{
	return 2*(r.width+r.height)
}

func main(){
	 // 创建一个接口类型的变量，并赋值一个实现了接口的具体类型的实例
	var sap Shape
	sap = Rectangle{2.2, 3.3}

	// 通过接口变量调用接口方法  
	fmt.Println("面积为 ",sap.Area())
	fmt.Println("周长为 ",sap.Perimeter())
}
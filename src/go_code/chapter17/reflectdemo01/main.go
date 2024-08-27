package main

import (
	"fmt"
	"reflect"
)

// 演示反射
func reflectTest01(b interface{}){

	//1.通过反射获取传入变量的type, kind, 值
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTpye=", rTyp)

	//2.获取reflect.value
	rVal := reflect.ValueOf(b)
	// n2 := 2 + rVal
	// fmt.Println("n2=",n2)
	fmt.Printf("rVal=%v, rVal type=%T\n", rVal,rVal)

	//3.将rVal转换interface{}
	iv := rVal.Interface()
	//通过断言转换
	num2 := iv.(int)
	fmt.Println("num2=", num2)
}

type Student struct {
	Name string
	Age int
}

//演示对结构体的反射
func reflectTest02(b interface{}){

	//1.通过反射获取传入变量的type, kind, 值
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTpye=", rTyp)

	//2.获取reflect.value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v, rVal type=%T\n", rVal,rVal)

	//3.将rVal转换interface{}
	iv := rVal.Interface()
	fmt.Printf("iv=%v, iv type=%T\n", iv,iv)

	// 获取变量对应的kind
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind1=%v, kind2=%v\n", kind1,kind2)

}

func main(){
	// var num  int = 100
	// reflectTest01(num)

	//定义结构体
	stu := Student{
		Name: "Tom",
		Age: 20,
	}
	reflectTest02(stu)
}
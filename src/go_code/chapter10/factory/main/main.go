package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

//因为student结构体首字母是小写，因此是只能在model使用//我们通过工厂模式来解决

func main(){
	var stu = model.NewStudent("tom",88)
	fmt.Println(*stu)
	fmt.Println("name=",stu.Name)
	fmt.Println("score=",stu.GetScore())
	fmt.Println()
}
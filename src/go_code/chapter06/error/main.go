package main

import (
	"errors"
	"fmt"
	"time"
)

func test() {
	defer func ()  {
		// recover()内置函数，可以捕获异常
		if err := recover();err != nil {
			fmt.Println("err=",err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=",res)

}

func readConf(name string) (err error){
	if name == "config.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

//自定义错误返回
func test02() {
	err := readConf("config.ini")
	if err != nil {
		// 如果读取文件错误，输出这个错误并终止程序
		panic(err)
	}
	fmt.Println("test02后面代码...")
}

func main() {
	// test()   //捕获错误实例
	test02()	//自定义错误实例
	for {
		fmt.Println("main后面代码...")
		time.Sleep(time.Second)
	}


}
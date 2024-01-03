package main

import (
	"fmt"
	"strconv"
)

func main(){
	// 统计字符串长度，汉字占3字节
	str := "hello北京"
	fmt.Println("str len=",len(str))

	// 字符串遍历
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0;i<len(r);i++{
		fmt.Printf("字符：%c\n",r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误",err)
	}else {
		fmt.Println("转换结果是：",n)
	}

	// 整数转字符串
	str = strconv.Itoa(123456)
	fmt.Printf("str=%v, typ=%T",str,str)

}
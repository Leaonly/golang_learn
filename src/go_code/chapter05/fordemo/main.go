package main

import "fmt"

func main() {
	/*
	// for 循环入门
	for i := 1; i <= 6;i++ {
		fmt.Println("hello, world!",i)
	}

	// 第二种写法
	j := 1
	for j <= 4 {
		fmt.Println("wow ", j)
		j++
	}

	// 第三种写法
	k := 1
	for {
		if k <= 3 {
			fmt.Println("ok lala~")
		} else {
			break
		}
		k++
	}
	*/

	//遍历字符串方法一，但如有中文会乱码，需切片转换
	var str string = "hello,world!你好!"
	str1 := []rune(str)
	for i := 0; i < len(str1);i++ {
		fmt.Printf("%c\n",str1[i])
	}

	//方法二 ，用for range，如有中文index 占3字节(3位)
	var str2 string = "hello,北京！"
	for index, val := range str2 {
		fmt.Printf("index=%d, val=%c\n",index,val)
	}
}
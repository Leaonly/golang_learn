package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串长度，汉字占3字节
	str := "hello北京"
	fmt.Println("str len=", len(str))

	// 字符串遍历
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符：%c\n", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果是：", n)
	}

	// 整数转字符串
	str = strconv.Itoa(123456)
	fmt.Printf("str=%v, typ=%T", str, str)

	// 字符串转byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// byte转字符串
	str = string([]byte{89, 29, 99})
	fmt.Printf("str=%v\n", str)

	// 10转2、8、16进制
	str = strconv.FormatInt(123, 2)
	fmt.Println("123的二进制是：", str)
	str = strconv.FormatInt(123, 16)
	fmt.Println("123的十六进制是：", str)

	// 查找子串是否在指定字符串中
	b := strings.Contains("seafood", "food")
	fmt.Printf("是否存在=%v\n", b) //true

	// 统计一个字符串有几个指定的子串
	num := strings.Count("chesse", "e")
	fmt.Printf("e有几个：%v\n", num)

	// 不区分大小写的字符串比较
	b = strings.EqualFold("abc", "ABC")
	fmt.Println("结果1为：", b) //true
	// ==是区分大小写
	fmt.Println("结果2位：", "abc" == "ABC") //false

	// 返回子串在字符串第一次出现的index
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index下标是=%v\n", index) //4 ,没有则为-1
	// 最后出现index
	index1 := strings.LastIndex("go_golang", "go")
	fmt.Printf("index1下标是=%v\n", index1)

	// 指定子串替换 n指定替换几个，-1全部替换
	str = strings.Replace("go go hello", "go", "啊哈", -1)
	fmt.Printf("str=%v\n", str)

	// 字符串分割, 为切片(动态数组)
	strArr := strings.Split("hello,world,你好", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strarr=%v\n", strArr)

	// 字符串大小写转换
	str = "goLang Hello"
	str = strings.ToLower(str)
	str1 := strings.ToUpper(str)
	fmt.Printf("str小写=%v\n", str)
	fmt.Printf("str大写=%v\n", str1)

	// 将字符串左右两边空格去掉
	str = strings.TrimSpace(" slej fjel fjie   ")
	fmt.Printf("str没空格=%v\n", str)

	// 将字符串左右两边指定字符串去掉
	str = strings.Trim("! helo@llo!", "!")
	fmt.Printf("str没叹号=%v\n", str)

	// 判断字符串是否以指定字符串开头
	b = strings.HasPrefix("ftp://192.168.0.1","ftp")
	fmt.Printf("b开头是ftp=%v\n", b)   //true
	b = strings.HasSuffix("xxx.jpg",".jpg")
	fmt.Printf("b结尾=%v\n", b)
}

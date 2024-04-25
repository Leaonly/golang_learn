package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	filepath := "d:/abc.txt"
	//O_TRUNC清空文件
	f,err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil{
		fmt.Println(err)
	}

	str := "你好 ,  wwww\r\n"
	writer := bufio.NewWriter(f)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	
	writer.Flush()
}
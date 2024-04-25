package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	filepath := "d:/abc.txt"
	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	str := "hello gradon.\n"
	writer := bufio.NewWriter(f)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}
	//bufio缓存写入文件
	writer.Flush()
}
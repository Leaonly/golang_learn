package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)  
	defer file.Close()
	
	//带缓存的读文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//输出内容
		fmt.Println(str)
	}
	fmt.Println("读取文件结束...")
}



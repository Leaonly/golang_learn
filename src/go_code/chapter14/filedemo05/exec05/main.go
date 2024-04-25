package main

import (
	//"bufio"
	"fmt"
	"os"
)

func main(){

	//读取abc文件内容到kkk文件
	f1 := "d:/abc.txt"
	f2 := "d:/kkk.txt"
	content, err:=os.ReadFile(f1)
	if err != nil{
		fmt.Println(err)
		return
	}
	f2o, err := os.OpenFile(f2, os.O_CREATE|os.O_WRONLY,0666)
	f2o.WriteString(string(content))
	defer f2o.Close()
	bu := os.IsNotExist(err)
	fmt.Println("文件是否存在",bu)
	
}
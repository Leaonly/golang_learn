package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "10.9.209.235:8888")
	if err != nil {
		fmt.Println("dial.conn err=", err)
		return
	}
	fmt.Println("conn sucess=", conn)
	defer conn.Close()

	//功能一： 客户端发送单行数据
	reader := bufio.NewReader(os.Stdin) //os.stdin代表终端输入

	for {
		//从终端读取一行，发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring err=", err)
		}
		//如果用户输入的是exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出..")
			break
		}
		//将line发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn write err=", err)
		}
		//fmt.Printf("客户端发送 %d 字节数据，退出\n", n)
	}
}

package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) (abc string) {
	//循环接收客户端发送的数据
	defer conn.Close()
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等待客户端conn发送消息
		//2.如果客户端没有write【发送】,协程就阻塞在这里
		fmt.Printf("服务端接收客户端%s信息: ", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("客户端退出 err=", err)
		// 	return
		// }
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		//3.现实客户端发送的内容到服务端
		content := string(buf[:n])
		fmt.Print(content)
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err= ", err)
		return
	}
	defer listen.Close() //延时关闭listen

	//循环等待连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(" Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v \n客户端IP=%v \n", conn, conn.RemoteAddr().String())
		}
		//准备一个协程
		go process(conn)

	}

}

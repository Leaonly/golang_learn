package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"io"
	"net"
	"time"
)

// 处理通信
func process(conn net.Conn) {
	defer conn.Close()

	//调用主控
	Processor := &Processor{
		Conn: conn,
	}
	err := Processor.Process2()
	if err != nil {
		if err == io.EOF {
			fmt.Println("Connetion closed by client")
		} else {
			fmt.Println("客户端和服务端通讯协程错误=err", err)
		}
	}
}

// 一个函数完成UserDao的初始化
func initDao() {
	//这里pool是一个redis.go全局变量
	model.MyUserDao = model.NewUserDao(pool)
	model.MySmsDao = model.NewSmsDao(pool)
}

func init() {
	//当服务器启动时，就初始化redis连接池
	//顺序注意先iniitPool,再initDao
	initPool("10.9.20.11:6379", "fNuu7Yhs5zQXa16WfMxQ", 16, 0, 300*time.Second)
	initDao()
}

func main() {
	//提示信息
	fmt.Println("服务器[重构]在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	//一旦监听成功，等待客户端连接
	fmt.Println("等待客户端连接....")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("net.Accept err=", err)
		}
		//连接成功，起协程通信
		go process(conn)
	}
}

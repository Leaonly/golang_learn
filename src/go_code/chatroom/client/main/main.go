package main

import (
	"fmt"
	"go_code/chatroom/client/processes"
	"os"
)

// 定义两个变量，一个表示id ，一个表示密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户选择
	var key int

	for {
		fmt.Println("------------欢迎登录多人聊天系统------------")
		fmt.Println("\t 1 登录聊天室")
		fmt.Println("\t 2 注册用户")
		fmt.Println("\t 3 退出系统")
		fmt.Println("\t 请选择(1-3): ")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			//登录
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录
			//1.创建一个userProcess实例
			up := &processes.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名(nickname):")
			fmt.Scanf("%s\n", &userName)
			//1.调用UserProcess实例的注册方法
			up := &processes.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入。")
			fmt.Println("")
		}
	}

}

/*
18.5.10 聊天的项目的扩展功能要求
1.实现私聊[点对点聊天]
2.如果一个登录用户离线，就把这个人从在线列表去掉
3.实现离线留言，在群聊时，如果某个用户没有在线，当登录后，可以接受离线的消息
*/

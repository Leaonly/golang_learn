package main

import (
	"fmt"
)

func main() {
	// 登录验证,3次机会
	var name string
	var passwd string
	var count int = 3
	for i:=0;i<3;i++ {
		fmt.Println("请输入用户名:")
		fmt.Scanln(&name)
		fmt.Println("请输入密码:")
		fmt.Scanln(&passwd)
		if name == "tom" && passwd == "888" {
			fmt.Println("登录成功！")
			break
		} else {
			count--
			fmt.Printf("用户名或密码错误，还有%v机会。\n",count)
		}
		if count == 0 {
			fmt.Println("账号已锁定，请30分钟后重试。")
		}
	}
}
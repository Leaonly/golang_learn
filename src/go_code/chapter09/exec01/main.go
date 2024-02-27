package main

import (
	"fmt"
)

/*
1)使用 map[string]map[string]sting 的 map 类型
2) key: 表示用户名，是唯一的，不可以重复
3)如果某个用户名存在，就将其密码修改”888888”，如果不存在就增加这个用户信息(包括昵称nickname和密码pwd)
4)编写一个函数 modifyUser(users map[string]map[string]sting,name string) 完成上述功能
*/
func modifyUser(users map[string]map[string]string,name string){
	if users[name] != nil {
		
		users[name]["passwd"] = "88888"
	} else {
		users[name] = make(map[string]string)
		users[name]["passwd"] = "88888"
		users[name]["nickname"] = "Tom Mask"
	}
}

func main(){
	users := make(map[string]map[string]string)
	modifyUser(users, "tom")
	modifyUser(users, "marry")
	fmt.Println(users)
}


package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	//定义小写结构体名字给其他域模块使用
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main(){
	monster := Monster{"悟空",2000,"七十二变"}
	//序列化为json
	jsonStr, err := json.Marshal(monster)
	if err != nil{
		fmt.Println("json序列化错误",err)
	}
	fmt.Println("jsonstr=",string(jsonStr))
}
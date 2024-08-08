package main

import (
	"encoding/json"
	"fmt"
)

// 定义结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Skill    string
}

// 反序列化json字符串为struct结构体
func unmarshalStruct() {
	str := "{\"Skill\":\"天马流星拳\",\"birthday\":\"100-1-1\",\"age\":20,\"name\":\"sam\"}"
	// 定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("反序列化后为: ", monster)
}

// 将map序列化
func testMap() string {
	// 定义一个map , 使用map需要make，或:=定义
	a := map[string]interface{}{}
	//a = make(map[string]interface{})
	a["name"] = "哪吒~~~~"
	a["age"] = 20
	a["address"] = "洪崖洞"

	// 将map a序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("map序列化: ",string(data))
	return string(data)
}

// 反序列化json字符串为map
func unmarshalMap() {
	//str := "{\"Skill\":\"哦豁\",\"birthday\":\"100-1-1\",\"age\":20,\"name\":\"map\"}"
	//如果json字符串是程序获取，则不需转义"处理
	str := testMap()
	// 定义一个map
	var a map[string]interface{}
	//反序列化map，不需要make,因为已被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("反序列化后为: ", a)
}

// 反序列化json字符串为切片
func unmarshalSlice() {
	str := "[{\"address\":\"美国\",\"age\":20,\"name\":\"tom\"}," +
		"{\"address\":\"墨西哥\",\"age\":20,\"name\":\"sam\"}]"
	// 定义一个slice
	var slice []map[string]interface{}
	//反序列化slice，不需要make,因为已被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("反序列化后为: ", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}

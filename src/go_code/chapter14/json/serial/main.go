package main

import (
	"fmt"
	"encoding/json"
)

// 定义结构体可以加tag，
type Monster struct{
	Name string 
	Age int
	Birthday string 
	Sal float64
	Skill string
}

func testStruct()  {
	monster := Monster{
		Name: "牛魔王",
		Age: 5000,
		Birthday: "100-1-1",
		Sal: 8000.0,
		Skill: "吃草",
	}

	//将monster 序列化
	data,err := json.Marshal(&monster)
	if err != nil{
		fmt.Println(err)
	}
	//输出序列化结果
	fmt.Println("struct序列化: ",string(data))
}

// 将map序列化
func testMap()  {
	// 定义一个map , 使用map需要make，或:=定义
	a:= map[string]interface{}{}
	//a = make(map[string]interface{})
	a["name"] = "哪吒"
	a["age"] = 20
	a["address"] = "洪崖洞"

	// 将map a序列化
	data,err := json.Marshal(a)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("map序列化: ",string(data))
}

//切片的json序列化
func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "tom"
	m1["age"] = 20
	m1["address"] = "美国"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "sam"
	m2["age"] = 20
	m2["address"] = "墨西哥"
	slice = append(slice, m2)

	//将切片序列化
	data,err := json.Marshal(slice)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("切片序列化: ",string(data))	
}

//对基本数据类型可以序列化，但没key value，没意义
func testFloat64()  {
	var num1 float64 = 234.24
	//将基本数据序列化
	data,err := json.Marshal(num1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("切片序列化: ",string(data))	
}

func main()  {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}

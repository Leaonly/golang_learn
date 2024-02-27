package main

import (
	"fmt"
)

func main(){
	// map的定义方法
	// 方式1
	var a map [string]string
	a = make(map[string]string, 10)
	a["n1"] = "宋江"
	a["n2"] = "吴用"
	a["n3"] = "武松"
	fmt.Println(a)

	// 方式2（推荐）
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 方式3
	heroes := map[string]string{
		"hero1" : "钢铁侠",
		"hero2" : "蜘蛛侠",
		"hero3" : "绿巨人",
	}
	fmt.Println("heroes=",heroes)

	// map的查找
	val, ok := cities["no1"]
	if ok {
		fmt.Println("有key，值为",val)
	} else {
		fmt.Println("没有key ...")
	}

	// 删除
	delete(cities,"no1")
	fmt.Println(cities)
	// 清空map，make新空间
	cities = make(map[string]string)
}


package main

import (
	"fmt"
)

func main(){
	// map遍历
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	for k, v := range cities {
		fmt.Printf("k=%v, v=%v\n",k,v)
	}

}
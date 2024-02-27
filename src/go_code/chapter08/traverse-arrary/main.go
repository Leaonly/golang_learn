package main

import (
	"fmt"
)

func main(){
	// 二维数组for 遍历
	var arr3 = [2][3]int{{1,2,3},{4,5,6}}

	for i := 0;i<len(arr3);i++{
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t",arr3[i][j])
		}
		fmt.Println()
	}

	// 二维数组for range遍历
	for _,v := range arr3{
		for _,v2 := range v{
			fmt.Printf("%v\t",v2)
		}
		fmt.Println()
	}
}
package main

import (
	//"errors"
	"fmt"
)

func judgeMonth(month int)int {
	a := 0
	if month >=1 && month <= 12 {
		day31 := []int{1, 3, 5, 7, 8, 10, 12}
		day30 := []int{4, 6, 9, 11}
		for _, value := range day31 {
			if month == value {
				a = 31
			}
		}
		for _, value := range day30 {
			if month == value {
				a = 30
			}
		}
		// if year % 4 == 0 &&  year % 100 != 0 || year % 400 == 0 {
		// 	fmt.Printf("%v是闰年!",year)
		// }
		if month == 2 {
			a = 28
		}
	}
	return a
}

func main (){
	var month int
	fmt.Println("请输入月份数字: ")
	fmt.Scanln(&month)
	day := judgeMonth(month)
	fmt.Println("该月份天数为: ",day)
}
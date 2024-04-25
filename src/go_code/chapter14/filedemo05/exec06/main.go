package main

import (
	"fmt"
	"os"
)

type TotalNum struct{
	Chnum int
	innum int
	spacenum int
	othernum int
}

func main(){
	filepath := "d:/test.txt"
	v, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var count TotalNum
	for _, b := range string(v) {
		switch {
		case b >= 'a' && b <= 'z':
			count.Chnum++
		case b >= 'A' && b <= 'Z':
			count.Chnum++
		case b >= '0' && b <= '9':
			count.innum++
		case b == ' ' || b == '\t':
			count.spacenum++
		default:
			count.othernum++
		}

	}
	fmt.Printf("计数，字符串%v个, 数字%v个, 空格%v个, 其他%v个",
	 count.Chnum,count.innum,count.spacenum,count.othernum)
}
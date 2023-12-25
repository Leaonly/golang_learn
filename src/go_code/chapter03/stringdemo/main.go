package main

import (
	"fmt"
	"strconv"
)

// str转基本数据类型
func main(){
	var str string = "true"
	var b bool
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T ; b=%v\n",b,b)

	var str2 string = "123456879"
	var n1 int64
	n1, _ = strconv.ParseInt(str2,10, 64)
	fmt.Printf("n1 type is %T ; n1=%v",n1,n1)
}
 
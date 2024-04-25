package main

import (
	"fmt"
	"flag"
)

func main(){
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","","用户名默认为空")
	flag.StringVar(&pwd,"pwd","","密码默认为空")
	flag.StringVar(&host,"h","","密码默认为空")
	flag.IntVar(&port,"p",3306,"端口默认为空")

	//调用前需转换&注册flag
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,port)
}
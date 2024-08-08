package main

import (
	"runtime"
	"fmt"
)

func main(){
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=",cpuNum)

	//设置使用多少个CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")

}

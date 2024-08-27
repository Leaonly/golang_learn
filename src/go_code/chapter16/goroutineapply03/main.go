package main
import (
	"fmt"

)

//向intchan放入1-8000个数
func putNum(intChan chan int){
	for i:=1;i<=8000;i++{
		intChan<-i
	}
	close(intChan)
}

//判断素数
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool){
	//var num int 
	for {
		num, ok := <- intChan
		if !ok {
			break
		}
		flag := true
		//判断num是不是素数
		for i := 2; i< num; i++{
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum 协程取不到数退出..")
	exitChan<-true
}

func main(){

	intChan := make(chan int, 1000)
	primeChan := make(chan int,2000)
	//标识退出管道,4个
	exitChan := make(chan bool, 4)

	//开启一个协程，向intchan放入1-8000个数
	go putNum(intChan)
	//开启四个协程取数放入primeChan
	for i := 0;i<4;i++{
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程判断流程完成
	go func () {
		for i :=0; i<4 ; i++{
			<-exitChan
		}
		//当exitChan取出四个结果，就可以关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan取出结果
	for {
		res, ok:= <- primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d \n",res)
	}

	fmt.Println("main线程退出...")
}
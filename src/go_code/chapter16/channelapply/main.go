package main
import (
	"fmt"

)

//write data
func writeData(intChan chan int)  {
	for i :=1; i<=50;i++{
		//放入数据
		intChan<-i
		fmt.Println("writeData ", i)
	}
	close(intChan)
}

//read data
func readData(intChan chan int,exitChan chan bool )  {
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData读取完数据则任务完成
	exitChan <- true
	close(exitChan)
}

func main(){

	//创建两个管道
	intChan := make(chan int,50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <- exitChan
		if !ok {
			break
		}
	}
}


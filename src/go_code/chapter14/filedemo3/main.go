package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.ReadFile("d:/test.txt")
	if err != nil {
		fmt.Println("err: ",err)
	}
	fmt.Println(string(file))
}
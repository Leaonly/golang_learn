package main

import (
	"fmt"
	"go_code/chapter11/object-encapsulation/model"
)

func main()  {

	model.SetUser("wo111w",1000,"flejl1")
	su := model.GetUser()
	fmt.Println(su)
}

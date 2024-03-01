package model

import "fmt"

type account struct {
	username string
	balance float64
	passwd string
}

func (acc *account)NewUser(u string,b float64,p string) *account {
	if 5 < len(u) && len(u) < 11  {
		acc.username = u
	} else {
		fmt.Println("输入账号长度有误")
		return nil
	}
	if 20 < b  {
		acc.balance = b
	} else {
		fmt.Println("输入余额有误")
		return nil
	}
	if len(p) == 6  {
		acc.passwd = p
	} else {
		fmt.Println("输入密码长度有误")
		return nil
	}
	return acc
}

func SetUser(u string,b float64,p string)*account  {
	var acc *account
	acc.NewUser(u,b,p)
	return acc
}


func GetUser()*account  {
	var acc *account
	return acc
}

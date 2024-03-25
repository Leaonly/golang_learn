package main

import (
	"fmt"
)

func main(){

	var key int
	loop := true
	// 余额
	balance := 1000.0
	//收支
	money := 0.0
	//说明
	note := ""
	//收支表头
	details := "收支\t账户金额\t收支金额\t说 明"
	for{
		fmt.Println("-------------家庭记账软件-------------")
		fmt.Println("1. 收支明细")
		fmt.Println("2. 登记收入")
		fmt.Println("3. 登记支出")
		fmt.Println("4. 退出软件")
		fmt.Println("请选择（1-4）：")
		fmt.Println()
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("-------------当前收支明细记录表-------------")
			fmt.Println(details)
			fmt.Println()
		case 2:
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance,money,note)
		case 3:
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			if money > balance {
				fmt.Println("余额不足...")
				break
			}
			balance -= money
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance,money,note)
		case 4:
			loop = false
		default:
			fmt.Println("请输入正确数字（1-4）...")
		}

		if !loop {
			break
		}
	}
	fmt.Println("退出软件...")
}



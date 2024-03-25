package utils

import (
	"fmt"
)

type FamilyAccount struct {
	key  string
	loop bool
	// 余额
	balance float64
	//收支
	money float64
	//说明
	note string
	//收支表头
	details string
}

type PersonalAccount struct{
	FamilyAccount
	name string
}

func NewFamilyAccount() *PersonalAccount {
	
	familyaccount := FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说 明",
	}
	personalAccount := PersonalAccount{
		FamilyAccount: familyaccount,
		name: "",
	}
		
	return &personalAccount
}

// 选择账户
func (this *PersonalAccount) choiceAccount() {
	fmt.Println("请输入账户: ")
	fmt.Scanln(&this.name)
}

// 打印收支明细方法
func (this *PersonalAccount) showDetails() {
	fmt.Println("-------------当前收支明细记录表-------------")
	fmt.Println(this.details)
	fmt.Println()
}

// 收入方法
func (this *PersonalAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)

}

// 支出方法
func (this *PersonalAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	if this.money > this.balance {
		fmt.Println("余额不足...")
		//break
	}
	this.balance -= this.money
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
}

// 退出方法
func (this *PersonalAccount) exit() {
	fmt.Println("确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，重新输入y/n")
	}
	if choice == "y" {
		this.loop = false
	}

}

// 给该结构体绑定方法
func (this *PersonalAccount) MainMenu() {
	for {
		fmt.Println("当前账户: ",this.name)
		fmt.Println("-------------家庭记账软件-------------")
		fmt.Println("0. 选择账户")
		fmt.Println("1. 收支明细")
		fmt.Println("2. 登记收入")
		fmt.Println("3. 登记支出")
		fmt.Println("4. 退出软件")
		fmt.Println("请选择（0-4）：")
		fmt.Println()
		fmt.Scanln(&this.key)
		switch this.key {
		case "0":
			this.choiceAccount()
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确数字（0-4）...")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("退出软件...")
}

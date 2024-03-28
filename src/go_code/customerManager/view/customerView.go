package main

import (
	"fmt"
	"go_code/customerManager/model"
	"go_code/customerManager/service"
)

type customerView struct {
	key  string
	loop bool //表示是否循环显示
	//增加service的字段
	customerService *service.CustomerService
}

// 显示所有客户信息
func (this *customerView) list() {
	customers := this.customerService.List()
	//显示
	fmt.Println("-------------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------------列表结尾-------------")
	fmt.Println()

}

func (this *customerView) add() {
	fmt.Println("-------------添加客户-------------")
	fmt.Println("姓名: ")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别: ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄: ")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话: ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱: ")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的customer实例, id号系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("添加完成!")
	} else {
		fmt.Println("添加失败...")
	}
}

// 得到id，删除id的客户
func (this *customerView) delete() {
	fmt.Println("-------------删除客户-------------")
	fmt.Println("请输入要删除的编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除
	}

	fmt.Println("确认是否删除(y/n): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用方法
		if this.customerService.Delete(id) {
			fmt.Println("删除成功！")
		} else {
			fmt.Println("删除失败....")
		}
	}
}

// 退出方法
func (this *customerView) exit() {
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

func (this *customerView)update()  {
	fmt.Println("请选择要修改的客户编号(-1退出): ")
	var id int
	fmt.Scanln(&id)
	fmt.Println("姓名: ")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别: ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄: ")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话: ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱: ")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的customer实例
	customer := model.NewCustomer(id,name, gender, age, phone, email)
	//调用
	if this.customerService.Update(customer) {
		fmt.Println("更新完成!")
	} else {
		fmt.Println("更新失败...")
	}

}

// 主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("-------------客户信息管理软件-------------")
		fmt.Println("0. 添加客户")
		fmt.Println("1. 修改客户")
		fmt.Println("2. 删除客户")
		fmt.Println("3. 客户列表")
		fmt.Println("4. 退出软件")
		fmt.Println("请选择（0-4）：")
		fmt.Println()
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "0":
			cv.add()
		case "1":
			cv.update()
		case "2":
			cv.delete()
		case "3":
			cv.list()
		case "4":
			cv.exit()
		default:
			fmt.Println("请输入正确数字（0-4）...")
		}

		if !cv.loop {
			break
		}
	}
	fmt.Println("退出软件...")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//初始化
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}

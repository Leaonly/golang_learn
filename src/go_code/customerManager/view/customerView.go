package main
import(
	"fmt"
	"go_code/customerManager/service"
)

type customerView struct{
	key string
	loop bool //表示是否循环显示
	//增加service的字段
	customerService *service.CustomerService
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

//主菜单
func (cv *customerView)mainMenu()  {
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
			cv.list()
		case "1":
			cv.list()
		case "2":
			cv.list()
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

//显示所有客户信息
func (this *customerView)list()  {
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

func main(){
	customerView := customerView{
		key: "",
		loop: true,
	}
	//初始化
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}

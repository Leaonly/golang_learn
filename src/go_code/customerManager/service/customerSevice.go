package service
import(
	"go_code/customerManager/model"
)


//对customer的操作
type CustomerService struct{
	customers []model.Customer
	//声明一个字段当前切片多少个客户
	customerNum int
}

func NewCustomerService()*CustomerService  {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",20,"112","abc@abc.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService)List()[]model.Customer {
	return this.customers
}

//添加客户到切片
func (this *CustomerService)Add(customer model.Customer) bool {
	//添加顺序id
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//根据id修改客户
func (this *CustomerService)Update(customer model.Customer) bool  {
	index := this.FindById(customer.Id)
	if index == -1 {
		return false
	}
	//修改切片对应id元素
	this.customers[index].Name = customer.Name
	this.customers[index].Age = customer.Age
	this.customers[index].Gender = customer.Gender
	this.customers[index].Phone = customer.Phone
	this.customers[index].Email = customer.Email

	return true
}

//根据id删除客户
func (this *CustomerService)Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	//从切片删除一个元素，切成index前后两块但不包括index
	this.customers = append(this.customers[:index], this.customers[index+1:]... )
	return true
}

//根据id查找切片下标，没有则返回-1
func (this *CustomerService)FindById(id int) int {
	index := -1
	//遍历切片
	for i:=0;i<len(this.customers);i++{
		if this.customers[i].Id == id {
			//找到
			index = i
		}
	}
	return index
}






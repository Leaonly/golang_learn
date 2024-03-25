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

func (this *CustomerService)List()[]model.Customer {
	return this.customers
}

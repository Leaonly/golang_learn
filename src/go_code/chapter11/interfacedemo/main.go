package main
import (
	"fmt"
)

type Usb interface{
	Start()
	Stop()
}

type Phone struct{

}
type Camera struct{

}

func (p Phone)Start()  {
	fmt.Println("手机开始工作...")
}
func (p Phone)Stop()  {
	fmt.Println("手机停止工作...")
}
func (p Camera)Start()  {
	fmt.Println("相机开始工作...")
}
func (p Camera)Stop()  {
	fmt.Println("相机停止工作...")
}

type Computer struct {

}

func (c Computer)Working(usb Usb)  {
	usb.Start()
	usb.Stop()
}

type Aint interface{
	Test01()
	Test02()
}

type Bint interface{
	Test01()
	Test03()
}

type Cint interface{
	Aint
	Bint
}




func main(){
	//创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

}
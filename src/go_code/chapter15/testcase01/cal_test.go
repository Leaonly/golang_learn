package cal 

import (
	//"fmt"
	"testing"
)

func TestAddUpper(t *testing.T)  {
	//调用
	res := addUpper(10)
	if res!=55{
		t.Fatalf("执行错误...期望值%v, 返回值%v",55,res)
	}
	t.Logf("执行成功")
}

func TestSubber(t *testing.T)  {
	//调用
	res := subber(10, 5)
	if res!=5{
		t.Fatalf("执行错误...期望值%v, 返回值%v",5,res)
	}
	t.Logf("执行成功")
}

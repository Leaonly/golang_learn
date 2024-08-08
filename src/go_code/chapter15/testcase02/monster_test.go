package monster

import (
	"testing"
)

func TestStore(t *testing.T){
	monster := &Monster{
		Name : "黑悟空",
		Age : 10,
		Skill : "三味真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为%v，实际为%v",true,res)
	}
	t.Logf("测试成功! ")
}

func TestReStore(t *testing.T)  {
	var monster Monster
	res := monster.ResStore()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为%v，实际为%v",true,res)
	}
	//进一步判断
	if monster.Name != "黑悟空"{
		t.Fatalf("monster.Store() 错误，希望为%v，实际为%v","黑悟空",monster.Name)

	}
	t.Logf("测试成功! ")
}

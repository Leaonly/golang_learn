package monster

import (
	"encoding/json"
	"fmt"
	"os"
)


type Monster struct {
	Name string
	Age int
	Skill string
}

//序列化struct保存到文件
func (this *Monster)Store() bool {
	//序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}

	//保存到文件
	filePath := "d:/Monster.ser"
	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}

//反序列化文件内容到struct
func (this *Monster)ResStore() bool {
	filepath := "d:/monster.ser"
	data,err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("readfile err =",err)
		return false
	}

	//对反序列化对象存到
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err =",err)
		return false
	}
	return true
}

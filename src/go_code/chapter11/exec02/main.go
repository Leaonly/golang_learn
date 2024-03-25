package main
import (
	"fmt"
	"sort"
	"math/rand"
)

//1.声明Hero结构体
type Hero struct{
	Name string
	Age int
}

//2.声明Hero结构体切片类型
type HeroSlice []Hero

//3.实现接口
func (hs HeroSlice)Len()int  {
	return len(hs)
}

func (hs HeroSlice)Less(i, j int)bool  {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice)Swap(i, j int)  {
	hs[i], hs[j] = hs[j], hs[i]
}


func main(){
	// 定义一个切片
	var intSlice = []int{0, -1, 31, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//
	var heroes HeroSlice
	for i:=0;i<10;i++{
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d",rand.Intn(100)),
			Age: rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	//排序前
	for _, v := range heroes{
		fmt.Println(v)
	}

	//排序后
	fmt.Println("----------------排序后--------------")
	sort.Sort(heroes)
	for _, v := range heroes{
		fmt.Println(v)
	}
}














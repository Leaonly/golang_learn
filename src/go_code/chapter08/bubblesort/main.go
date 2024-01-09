package main

import (
	"fmt"
)

// 冒泡排序
func BubbleSort(arr *[]int)  {
	fmt.Println("排序前arr=",(*arr))
	for j:=0;j<len(*arr)-1;j++{
		for i:=0;i<len(*arr)-1-j;i++{
			if (*arr)[i]>(*arr)[i+1] {
				(*arr)[i],(*arr)[i+1] = (*arr)[i+1],(*arr)[i]
			}
		}
	}
}

//冒泡排序2
func BubbleSort2(arr *[]int)  {
	fmt.Println("排序前arr=",(*arr))
	for j := 0; j < len(*arr)-1; j++ {
		for i:=0;i<len(*arr)-1-j;i++{
			if (*arr)[i] > (*arr)[i+1] {
				(*arr)[i],(*arr)[i+1] = (*arr)[i+1],(*arr)[i]
			}
		}
	}
}

func main(){
	arr := []int{24,10,48,9,88,2,49}
	BubbleSort2(&arr)
	fmt.Println("排序后arr=",(arr))
}
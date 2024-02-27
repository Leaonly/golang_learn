package main

import (
	"fmt"
)

func BinaryFind(arr *[]int, leftIndex int, rightIndex int,findVal int)  {
	
	//判断左标是否大于右标，则为找不到
	if leftIndex > rightIndex {
		fmt.Println("找不到...")
		return
	}

	// 找到中间下标值等于findVal
	middle := (rightIndex + leftIndex)/2
	if (*arr)[middle] > findVal {
		BinaryFind(arr,leftIndex,middle-1,findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr,middle+1,rightIndex,findVal)
	} else {
		fmt.Printf("找到了,下标为%v\n",middle)
	}
}

func main(){
	arr := []int{1,29,33,66,199,505,5000,20200}
	BinaryFind(&arr,0,7,505)
}
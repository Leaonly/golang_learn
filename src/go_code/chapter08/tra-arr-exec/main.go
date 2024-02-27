package main

import (
	"fmt"
	"math/rand"
)

func BinaryFind2(arr3 *[10]int, leftIndex int, rightIndex int, judgeNum int){

	if leftIndex > rightIndex {
		fmt.Println("没找到...")
		return
	}

	middle := (leftIndex + rightIndex)/2
	if (*arr3)[middle] > judgeNum {
		BinaryFind2(arr3, leftIndex, middle -1, judgeNum)
	} else if (*arr3)[middle] < judgeNum {
		BinaryFind2(arr3, middle + 1, rightIndex, judgeNum)
	} else {
		fmt.Println("找到了,下标为",middle)
	}
}

func main() {
	/*
	// 题目1：随机生成10个整数(1 100 的范围)保存到数组，
	//并倒序打印以及求平均值、求最大值和最大值的下标、并查找里面是否有 55
	arr1 := [10]int{}
	for i := 0; i < 10; i++ {
		num := rand.Intn(100) + 1
		arr1[i] = num
	}
	fmt.Println(arr1)

	for i := 9; i >= 0; i-- {
		fmt.Printf("%v ",arr1[i])
	}
	fmt.Println()
	// 求平均
	sum := 0
	maxNum := arr1[0]
	minNum := arr1[0]
	for _,v := range arr1 {
		sum += v
		if v > maxNum {
			maxNum = v
		}
		if v < minNum{
			minNum = v
		}
	}
	fmt.Println("平均值为：",float64(sum/len(arr1)))

	for i,v := range arr1{
		if maxNum == v {
			fmt.Printf("最大值为%v,下标为%v\n",maxNum,i)
		}
		if minNum == v {
			fmt.Printf("最小值为%v,下标为%v\n",minNum,i)
		}
		if v == 55 {
			fmt.Println("数组里有55!")
		}
	}
	
	
	// 题目2 
	arr2 := [3][4]int{}
	for i,_ := range arr2{
		for j,_ := range arr2[i]{
			num := rand.Intn(100)
			arr2[i][j] = num
			fmt.Printf("%v ",arr2[i][j])
		}
		fmt.Println()
	}
*/
	// 题目3：随机生成 10个整数(1-100 之间)，使用冒泡排序法进行排序，
	// 然后使用二分查找法，查找是否有90 这个数，
	// 并显示下标，如果没有则提示“找不到该数”
	arr3 := [10]int{}
	for i,_ := range arr3{
		num := rand.Intn(100) + 1
		arr3[i] = num
	}
	fmt.Println("数组为",arr3)
	for i := 0; i < len(arr3)-1; i++ {
		for j := 0; j < len(arr3)-1-i; j++ {
			if arr3[j] > arr3[j+1] {
				arr3[j], arr3[j+1] = arr3[j+1], arr3[j]
			}
		}
	}
	fmt.Println("排序后为",arr3)

	BinaryFind2(&arr3,0,9,90)

}


package main

import "fmt"

func TestArr() {
	var arr [10]int //array宣告必須宣告容量
	fmt.Println("arr : ", arr)
}

func TestArr2() {
	var arr1 [5]int
	arr2 := [5]int{0}

	arr3 := [...]int{0, 0, 0, 0, 1}
	arr4 := [9]int{0, 0, 9, 0, 0}

	fmt.Println("arr1 : ", arr1)
	fmt.Println("arr2 : ", arr2)
	fmt.Println("arr3 : ", arr3)
	fmt.Println("arr4 : ", arr4)
	fmt.Println("arr1 == arr2 : ", arr1 == arr2)
	fmt.Println("arr1 == arr3 : ", arr1 == arr3) // 同型別會判斷內容物是否完全相同
	//fmt.Println("arr1 == arr4 : ", arr1 == arr4) // 不同型別([5]int [9]int)不能拿來做比較，會發生錯誤
}

func TestArr3() {
	arr := [5]int{1: 5} // 初始化index 1的直等於5
	fmt.Println("arr : ", arr)
	fmt.Println("arr[1] : ", arr[1])

	// array 賦值
	arr[0] = 10
	fmt.Println("arr[0] : ", arr[0])
}

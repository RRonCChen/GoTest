package main

import "fmt"

func main() {
	fmt.Println(testReturn("這是"))
	testVariadic("不固定變數傳入", 1, 2, 3, 4, 5)

	i := []int{2, 4, 6, 8, 10}
	testVariadic("slice傳入", i...) //若要用slice傳入必須使用unpack operator
}

func testReturn(s string) (ss string) {
	ss = s + "測試"
	return //具名的return,自動return ss
}

func testVariadic(s string, i ...int) {
	fmt.Print(s)
	fmt.Println(i) //不固定數量變數,會被轉換成slice傳入
}

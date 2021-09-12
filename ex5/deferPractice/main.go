package main

import "fmt"

func main() {
	i := 10
	defer deferTest(i) //會在父函式中最後一個執行
	i += 10
	defer deferTest(i) //FILO的順序執行,值會是當下輸入的狀態
	fmt.Println("main開始")
	fmt.Println("main結束")
}

func deferTest(i int) {
	fmt.Println("i : ", i)
}

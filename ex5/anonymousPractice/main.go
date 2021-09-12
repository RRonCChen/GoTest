package main

import "fmt"

func main() {

	func() {
		fmt.Println("匿名函式宣告直接呼叫練習")
	}()

	f1 := func(meg string) {
		fmt.Println(meg)
	}
	f1("測試匿名函式指派")
}

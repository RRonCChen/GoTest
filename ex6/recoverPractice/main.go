package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("main結束")
}

func a() {
	b()
	fmt.Println("a結束")
}

func b() {

	defer func() {
		if r := recover(); r != nil { // r 回傳panic輸入的interface{}
			fmt.Println("b()發生recover : ", r) //確保發生panic後程式能夠跑完
		}
	}()

	panic(errors.New("b()發生panic")) //輸入的interface會是recover回傳值

	fmt.Println("b結束") //此行不會執行到
}

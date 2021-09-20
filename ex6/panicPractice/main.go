package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("呼叫前")
	defer deferTest()
	panicTest()
	fmt.Println("呼叫後") //不會被呼叫
}

func deferTest() {
	fmt.Println("defer function")
}

func panicTest() {
	fmt.Println("panic function")
	defer fmt.Println("panic defer")
	panic(errors.New("發生主動panic"))
}

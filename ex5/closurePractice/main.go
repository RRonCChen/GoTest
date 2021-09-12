package main

import "fmt"

func main() {
	i := 0
	increment := func() int {
		i++
		return i
	}
	fmt.Println(increment())
	fmt.Println(increment())
	i += 100
	fmt.Println(increment())

	fmt.Println("閉包測試01")
	inc := incrementor()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println("閉包測試02")
	fmt.Println(incrementor()())
	fmt.Println(incrementor()())
	fmt.Println(incrementor()())
}

func incrementor() func() int {
	i := 0
	//閉包會記得父函數外部變數
	return func() int {
		i++
		return i
	}
}

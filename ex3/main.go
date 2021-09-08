package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	testMathBig()
	testString()
	testRune()
	testNil()
}

func testMathBig() {

	intA := math.MaxInt64
	intA += 2 //發生越界繞回
	fmt.Println(intA)

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(2)) //正確得加上去
	fmt.Println(bigA)
}

func testString() {

	str1 := "你好"

	for i := 0; i < len(str1); i++ {
		fmt.Print(string(str1[i])) //單一byte讀取會發生錯誤（UTF-8容量可能超過一個byte）
	}

	fmt.Println(" ")

	runes := []rune(str1) //可包含 1-4個byte的容量符合UTF-8
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]))
	}

}

func testRune() {
	str := "你好"
	runes := []rune(str)

	//擷取string 字串個數時 轉為rune slice會是正確的個數
	fmt.Println(len(str))   // 會列出str得byte個數
	fmt.Println(len(runes)) //字元個數

}

func testNil() {
	var message *string = new(string)

	*message = "測試"

	if message == nil {
		fmt.Println("message為空值")
		return
	}
	fmt.Println(*message)
}

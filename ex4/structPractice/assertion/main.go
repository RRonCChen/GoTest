package main

import (
	"errors"
	"fmt"
)

func main() {
	res, _ := doubler(2)
	fmt.Println("2 : ", res)

	assertSwitch(int16(5))
	assertSwitch(5)

}

func doubler(v interface{}) (string, error) {
	if i, ok := v.(int); ok { //型別斷言檢查型別   value , boolean := element.(型別)
		return fmt.Sprint(i * 2), nil
	}
	return "", errors.New("傳入未支援的值")
}

func assertSwitch(v interface{}) {
	//使用switch來判斷型別斷言
	switch v.(type) {
	case string:
		fmt.Println("v is string")
	case int:
		fmt.Println("v is int")
	default:
		fmt.Println("v is other type")
	}
}

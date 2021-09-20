package main

import (
	"errors"
	"fmt"
)

var (
	ErrNumberRange = errors.New("超過數字大小")
)

func main() {
	num := 100
	if err := checkNumber(num); err != nil {
		fmt.Println(err)
	}

	if err := checkNumber2(num); err != nil {
		fmt.Println(err)
	}
}

func checkNumber(num int) error {
	if num < 10 || num > 0 {
		return ErrNumberRange
	}
	return nil
}

func checkNumber2(num int) error {
	if num < 10 || num > 0 {
		return fmt.Errorf("檢驗數字錯誤 ： %w", ErrNumberRange) //wrap錯誤, error chain
	}
	return nil
}

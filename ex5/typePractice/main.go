package main

import "fmt"

type calc func(a, b int) string

func main() {
	calculator(add, 1, 2)
	calculator(minus, 2, 1)
}

func add(a, b int) string {
	result := a + b
	return fmt.Sprintf("%d + %d = %d", a, b, result)
}

func minus(a, b int) string {
	result := a - b
	return fmt.Sprintf("%d - %d = %d", a, b, result)
}

func calculator(f calc, a, b int) {
	fmt.Println(f(a, b))
}

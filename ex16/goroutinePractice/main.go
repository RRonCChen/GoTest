package main

import (
	"fmt"
	"time"
)

func main() {
	var sum1, sum2 int

	go func() {
		sum1 = sum(1, 10)
	}()
	sum2 = sum(1, 100)
	time.Sleep(time.Second)

	fmt.Println(sum1, sum2)
}

func sum(strat, end int) int {
	result := 0
	for i := 0; i <= end; i++ {
		result += i
	}
	return result
}

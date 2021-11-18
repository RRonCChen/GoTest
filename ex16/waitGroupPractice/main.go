package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum1, sum2 int

	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(2) // 告訴waitgroup 要等待幾個goruntine結束(若多宣告會造成deadLock)
	go sum(1, 10, waitGroup, &sum1)
	go sum(1, 100, waitGroup, &sum2)
	waitGroup.Wait() // 等待goruntine執行
	fmt.Println(sum1, sum2)

}

func sum(start, end int, waitGroup *sync.WaitGroup, result *int) {
	*result = 0
	for i := 0; i <= end; i++ {
		*result += i
	}
	if waitGroup != nil {
		waitGroup.Done()
	}
}

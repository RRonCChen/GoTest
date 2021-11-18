package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go sum(1, 25, wg, &s1)
	go sum(26, 50, wg, &s1)
	go sum(51, 75, wg, &s1)
	go sum(76, 100, wg, &s1)
	wg.Wait()

	fmt.Println("s1 = ", s1)
}

func sum(from, to int, wg *sync.WaitGroup, res *int32) {
	for i := from; i <= to; i++ {
		atomic.AddInt32(res, int32(i)) // sync.atomic.AddInt系列，確保變數一次只能有一個來源寫入（原子性）
	}
	wg.Done()
}

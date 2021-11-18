package main

import (
	"fmt"
	"sync"
)

func main() {
	res := 0
	mut := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go sum(1, 25, wg, mut, &res)
	go sum(26, 50, wg, mut, &res)
	go sum(51, 75, wg, mut, &res)
	go sum(76, 100, wg, mut, &res)
	wg.Wait()
	fmt.Println("result : {0}", res)
}

func sum(from, to int, wg *sync.WaitGroup, mut *sync.Mutex, res *int) {
	for i := from; i <= to; i++ {
		mut.Lock()
		*res += i
		mut.Unlock()
	}
	if wg != nil {
		wg.Done()
	}
}

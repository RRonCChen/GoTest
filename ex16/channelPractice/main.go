package main

import (
	"fmt"
	"time"
)

func main() {
	// chanenl練習1
	fmt.Println("ch1 -------------------->")
	ch := make(chan int, 1) // 初始化通道,緩衝區大小 = 1,若使用無緩衝區在main會造成deadlock
	defer close(ch)         // 關閉通道,關閉後不可再塞入值,讀取的話會取道該型別的零值
	ch <- 1                 // 將1傳入通道
	i, ok := <-ch           // 從通道取出
	if ok {
		fmt.Printf("i : %d\n", i)
	}

	// chanenl練習2
	fmt.Println("ch2 -------------------->")
	ch2 := make(chan string) //建立無緩衝區的字串通道
	go func() {
		ch2 <- "hello"
	}()
	fmt.Println(<-ch2)

	// chanenl練習3
	fmt.Println("ch3 -------------------->")
	ch3 := make(chan string)
	go func() {
		msg := <-ch3
		ch3 <- fmt.Sprintf("接收到訊息了 ： %s", msg)
		ch3 <- "hello main from greet"
	}()
	ch3 <- "hello greet from main"
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)

	// channel練習4
	fmt.Println("ch4 -------------------->")
	ch4 := make(chan int)
	finalSum := 0
	go sum(1, 25, ch4)
	go sum(26, 50, ch4)
	for c := 0; c < 50; c++ {
		v := <-ch4
		finalSum += v
	}
	fmt.Printf("sum : %d\n", finalSum)

	// channel練習5
	fmt.Println("ch5 -------------------->")
	ch5 := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch5 <- i
		}
		close(ch5)
	}()
	for {
		v, ok := <-ch5
		if !ok { // 判斷channel是否close
			break
		}
		fmt.Println(v)
	}

	// channel練習6
	fmt.Println("ch6 -------------------->")
	ch6 := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			ch6 <- i
		}
		close(ch6)
	}()
	for i := range ch6 { //range可以巡迴channel直到close
		fmt.Println(i)
	}

	// channel練習7
	fmt.Println("ch7 -------------------->")
	ch7 := make(chan int)
	go func() {
		time.Sleep(1000)
		ch7 <- 100
	}()
	for {
		select { // 用select來解決阻塞問題（ 會等到沒有阻塞情況時(channel 內有資料)才會執行）
		case i := <-ch7:
			fmt.Printf("ch7 : %d\n", i)
			return
		default:
			fmt.Println("運算中...")
		}
	}
}

func sum(from, to int, out chan int) {
	for i := from; i <= to; i++ {
		out <- i
		time.Sleep(time.Microsecond * 2)
	}
}

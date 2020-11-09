package main

import (
	"fmt"
	"time"
)

func testDeadLock(c chan int) {
	fmt.Println(<-c)
	//for{
	//	fmt.Println(<-c)
	//}
}

func main() {
	//c := make(chan int, 1)
	//
	//// make(chan int)，开辟的通道是一种无缓冲通道，所以当对这个缓冲通道写的时候，会一直阻塞等到某个协程对这个缓冲通道读
	//c <- 10
	//// fatal error: all goroutines are asleep - deadlock!
	//go testDeadLock(c)
	//
	//time.Sleep(time.Millisecond)

	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
		}(i)
	}

	fmt.Println(<-ch)           // 输出第一个发送的值
	close(ch)                   // 不能关闭，还有其他的 sender
	time.Sleep(2 * time.Second) // 模拟做其他的操作
}

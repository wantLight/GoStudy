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
	c := make(chan int, 1)

	// make(chan int)，开辟的通道是一种无缓冲通道，所以当对这个缓冲通道写的时候，会一直阻塞等到某个协程对这个缓冲通道读
	c <- 10
	// fatal error: all goroutines are asleep - deadlock!
	go testDeadLock(c)

	time.Sleep(time.Millisecond)
}

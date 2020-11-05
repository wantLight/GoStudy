package main

import (
	"fmt"
	"time"
)

/**
我们上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢，
Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，
当多个channel都准备好的时候，select是随机的选择一个执行的。
*/
func fibonacci(c, wang, quit chan int) {
	x, y := 1, 1
	for {
		select {
		// 只有当监听的channel中有发送或接收可以进行时才会运行
		case c <- x:
			fmt.Printf("c%v\n", x)
			x, y = y, x+y
		//case wang <- y:
		//	fmt.Printf("wang%v\n", y)
		//	y = 233
		case <-quit:
			fmt.Println("quit")
			return
		case <-time.After(5 * time.Second):
			println("timeout")
			wang <- 666
			return

		}
	}
}

func main() {
	c := make(chan int)
	wang := make(chan int)
	quit := make(chan int)

	//func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//		fmt.Println(<-wang)
	//	}
	//	//quit <- 0
	//}()

	go fibonacci(c, wang, quit)
	fmt.Println(<-wang)

	//c := make(chan int)
	//o := make(chan bool)
	//go func() {
	//	for {
	//		select {
	//		case v := <- c:
	//			println(v)
	//		case <- time.After(5 * time.Second):
	//			println("timeout")
	//			o <- true
	//			break
	//		}
	//	}
	//}()
	//<- o
}

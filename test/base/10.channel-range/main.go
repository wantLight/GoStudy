package main

import "fmt"

// 摘自：https://www.yiibai.com/go/golang-range-over-channels.html
// 摘自：https://gobyexample.com/range-over-channels

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}

}

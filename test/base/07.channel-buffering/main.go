package main

import "fmt"

// 摘自：https://www.yiibai.com/go/golang-channel-buffering.html
// 摘自：https://gobyexample.com/channel-buffering

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	//messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	//fmt.Println(<-messages)
}

package main

import "fmt"

// 摘自：https://www.yiibai.com/go/golang-channels.html
// 摘自：https://gobyexample.com/channels

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	// 创建无缓冲的管道，同 `messages := make(chan string, 0)`。
	messages := make(chan string)

	// Send a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The `<-channel` syntax receives a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 摘自：https://gobyexample.com/signals

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("awaiting signal.")
	sig := <-sigs
	fmt.Println(sig, "signal comes.")
	fmt.Println("exiting.")
}

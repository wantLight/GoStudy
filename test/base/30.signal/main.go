package main

import (
	"fmt"
	"os"
	"os/signal"
)

// 摘自：https://gobyexample.com/signals

func main() {

	sigs := make(chan os.Signal, 1)
	//notify用于监听信号
	//参数1表示接收信号的channel
	//参数2及后面的表示要监听的信号
	//os.Interrupt 表示中断
	//os.Kill 杀死退出进程
	signal.Notify(sigs, os.Interrupt, os.Kill)

	fmt.Println("awaiting signal.")
	sig := <-sigs
	fmt.Println(sig, "signal comes.")
	fmt.Println("exiting.")
}

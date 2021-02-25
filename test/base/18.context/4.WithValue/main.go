package main

import (
	"context"
	"fmt"
	"time"
)

// 摘自：https://my.oschina.net/mrrdai/blog/1838998

func someHandler() {
	ctx := context.WithValue(context.Background(), "a", "b")
	go doStuff(ctx)
}

func doStuff(ctx context.Context) {
	fmt.Println(ctx.Value("a"))
}

func main() {
	someHandler()

	// 等待子协程结束。
	time.Sleep(1 * time.Second)
}

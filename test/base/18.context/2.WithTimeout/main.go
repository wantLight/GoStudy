package main

import (
	"context"
	"fmt"
	"time"
)

// 摘自：https://my.oschina.net/mrrdai/blog/1838998

func someHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go doStuff(ctx)

	// 5秒后取消doStuff。
	time.Sleep(5 * time.Second)
	cancel()
}

// 每1秒work一下，同时会判断ctx是否被取消了，如果是就退出。
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("work")
		}
	}
}

func main() {
	someHandler()

	// 等待子协程结束。
	time.Sleep(1 * time.Second)
}

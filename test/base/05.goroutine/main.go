package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 摘自：GO 语言简介（下）— 特性 https://coolshell.cn/articles/8489.html

func print(name string, delay time.Duration) {

	t0 := time.Now()
	fmt.Println(name, "start at", t0)

	for i := 0; i < 1000; i++ {
		fmt.Println(name, i)
	}

	t1 := time.Now()
	fmt.Println(name, "end at", t1, "lasted", t1.Sub(t0))
}

func main() {

	// GOMAXPROCS最好不要超过NumCPU，不然CPU之间切换也是浪费时间。
	runtime.GOMAXPROCS(3)

	// 生成随机种子。
	rand.Seed(time.Now().Unix())

	var name string
	for i := 0; i < 3; i++ {

		// 生成ID
		name = fmt.Sprintf("go_%02d", i)

		// 生成随机等待时间，从0-4秒。
		go print(name, time.Duration(rand.Intn(5))*time.Second)
	}

	// 让主进程停住，不然主进程退了，goroutine也就退了。
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

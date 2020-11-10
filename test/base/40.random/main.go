package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 将时间戳设置成种子数。
	rand.Seed(time.Now().UnixNano())

	fmt.Println("生成10个[0-100)之间的随机数：")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	fmt.Println("生成10个[50-150)之间的随机数：")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100) + 50)
	}
}

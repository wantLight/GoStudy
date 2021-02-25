package main

import (
	"fmt"
	"time"
)

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func myRange(mychan chan int) {
	for v := range mychan {
		fmt.Println(v)
	}
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum1(s[:len(s)/2], c)
	//go sum1(s[len(s)/2:], c)
	//go sum1([]int{1, 2, 3}, c)
	//// 从 c 中接收
	//x, y, z := <-c, <-c, <-c
	//fmt.Println(x, y, z, x+y)

	samChan := make(chan int, 3)
	go myRange(samChan)
	samChan <- 8
	samChan <- 7
	samChan <- 6
	samChan <- 5
	samChan <- 4
	// 在生产者的地方关闭channel
	// 如果向此channel写数据的goroutine退出时，系统检测到这种情况后会panic，否则range将会永久阻塞。
	close(samChan)
	time.Sleep(time.Second * 10) // 正常

}

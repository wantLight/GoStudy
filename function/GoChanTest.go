package main

import "fmt"

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum1(s[:len(s)/2], c)
	go sum1(s[len(s)/2:], c)
	go sum1([]int{1, 2, 3}, c)
	x, y, z := <-c, <-c, <-c // 从 c 中接收

	fmt.Println(x, y, z, x+y)
}

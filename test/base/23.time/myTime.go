package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now())
	//<-time.After(time.Second*1)
	//fmt.Println(time.Now())
	//
	//log.Println("AfterFuncDemo start: ", time.Now())
	//// time.AfterFunc()是异步执行的，所以需要在函数最后sleep等待指定的协程退出，否 则可能函数结束时协程还未执行。
	//time.AfterFunc(time.Second*2, func() {
	//	log.Println("AfterFuncDemo end: ", time.Now())
	//})
	//time.Sleep(3 * time.Second) // 等待协程退出

	TickerLaunch()
}

func GetNewPassenger() string {
	return "xiaohu"
}

// TickerLaunch用于演示ticker聚合任务用法
func TickerLaunch() {
	ticker := time.NewTicker(1 * time.Minute)
	maxPassenger := 10 // 每车最大装载人数
	passengers := make([]string, 0, maxPassenger)

	for {
		passenger := "xiaohu" // 获取一个新乘客
		if passenger != "" {
			passengers = append(passengers, passenger)
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(2 * time.Second)
		}

		select {
		case <-ticker.C: // 时间到，发车
			fmt.Println(" 时间到，发车")
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger { // 时间没到，车已座满，发车
				fmt.Println("时间没到，车已座满，发车")
				passengers = []string{}
			}
		}
	}
}

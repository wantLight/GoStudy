package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	_, err := os.Open("filename.ext")
	// 类似于os.Open函数，标准包中所有可能出错的API都会返回一个error变量，以方便错误处理
	// 可以通过errors.New把一个字符串转化为errorString，以得到一个满足接口error的对象
	if err != nil {
		log.Fatal(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// implementation
}

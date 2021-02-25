package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
)

// 参考：https://blog.csdn.net/kongxx/article/details/76167000

func main() {

	dataBytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("Something went wrong. err=[%+v].\n", err)
		debug.PrintStack()
		os.Exit(1)
	}

	fmt.Printf("Read data successfully. data=[%+v].", string(dataBytes))
}

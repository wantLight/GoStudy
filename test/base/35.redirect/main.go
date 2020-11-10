package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"syscall"
)

func main() {

	// 打开文件。
	f, err := os.OpenFile("/tmp/redirect.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("Something went wrong. err=[%+v].\n", err)
		debug.PrintStack()
		os.Exit(1)
	}

	// 重定向。
	RedirectOutput(f)

	// 测试重定向。
	fmt.Println("Println Output to file [ /tmp/redirect.log ], not to Stdout.")
	fmt.Printf("Printf Output to file [ /tmp/redirect.log ], not to Stdout.\n")
	fmt.Print("Print Output to file [ /tmp/redirect.log ], not to Stdout.\n")
	panic("Panic Output to file [ /tmp/redirect.log ], not to Stdout.\n")
}

// RedirectOutput 重定向Stdout和Stderr到指定文件。
// 参考：https://www.cnblogs.com/ghj1976/p/4276390.html
// 参考：https://github.com/golang/go/issues/325
func RedirectOutput(file *os.File) {

	// 注意：可以单独使用高层API，或者单独使用底层API，或者同时使用高层和底层API进行重定向。同时使用高层和底层API时，也不会发生一条消息输出两次的问题。

	// 使用高层API进行重定向。支持fmt，不支持panic。
	os.Stdout = file
	os.Stderr = file

	// 使用底层API进行重定向。支持fmt和panic。
	// syscall.Dup2仅是linux/OSX的API，windows上没有。
	// syscall.Dup2 is a linux/OSX only thing. there's no windows equivalent。
	syscall.Dup2(int(file.Fd()), syscall.Stdout)
	syscall.Dup2(int(file.Fd()), syscall.Stderr)
}

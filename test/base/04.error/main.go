package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// 摘自：Go 语言简介（上）— 语法 http://coolshell.cn/articles/8460.html

// ////////////////////////////////////////////////////////////////////
// 错误处理：函数错误返回可能是C/C++时最让人纠结的东西的，Go的多值返回可以让我们更容易的返回错误，
// 其可以在返回一个常规的返回值之外，还能轻易地返回一个详细的错误描述。通常情况下，错误的类型是error，它有一个内建的接口。
// ////////////////////////////////////////////////////////////////////
// 自定义的出错结构
type MyError struct {
	arg    int
	errMsg string
}

// 实现Error接口
func (e *MyError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
}

// 其可以在返回一个常规的返回值之外，还能轻易地返回一个详细的错误描述
func square(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("Bad Arguments - negtive!")
	} else if arg > 256 {
		return -1, &MyError{arg, "Bad Arguments - too large!"}
	}
	return arg * arg, nil
}

func testError() {
	for _, i := range []int{-1, 4, 1000} {
		if r, e := square(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("success:", r)
		}
	}
}

// ////////////////////////////////////////////////////////////////////
// 错误处理：defer注册执行的方法，在CopyFile方法返回后自动调用资源释放，这种解决方式还是比较优雅的
// ////////////////////////////////////////////////////////////////////
func copyFile(dstName, srcName string) (written int64, err error) {

	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	testError()
}

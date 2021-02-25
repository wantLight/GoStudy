package main

import (
	"fmt"
	"sync"

	"github.com/xdc0209/go-base/utils/time_util"
)

// 摘自：https://www.cnblogs.com/sunsky303/p/9706210.html
//
// sync.Pool：临时对象池。
// 临时对象池比较适合用来存储一些临时且状态无关的数据，但是不适合用来存储数据库连接的实例，因为存入临时对象池重的值有可能会在垃圾回收时被删除掉，这违反了数据库连接池建立的初衷。
// 临时对象池主要作用是减少GC，提高性能，所以别被别的池子带跑了，golang里的sync.Pool就是GC优化的。
// 还有一点需要注意，根据实际情况需要在Put前或Get后重置临时对象，以免产生意想不到的问题，可参考fmt/print.go:Sprintf中的实现。
//
// 案例1：gin 中的 Context pool
// github.com/gin-gonic/gin/gin.go:ServeHTTP
//
// 案例2：fmt 中的 Printer pool
// fmt/print.go:Sprintf
// fmt/print.go:newPrinter
// fmt/print.go:free

var bytesPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

func main() {

	// without pool:
	startTime1 := time_util.Milliseconds()
	var obj1 []byte
	for i := 0; i < 100000000; i++ {
		obj := make([]byte, 1024)
		obj1 = obj
	}
	endTime1 := time_util.Milliseconds()

	// with pool:
	startTime2 := time_util.Milliseconds()
	var obj2 []byte
	for j := 0; j < 100000000; j++ {
		obj := bytesPool.Get().([]byte)
		obj2 = obj
		bytesPool.Put(obj)
	}
	endTime2 := time_util.Milliseconds()

	fmt.Println(obj1)
	fmt.Println(obj2)
	fmt.Println("without pool: ", endTime1-startTime1, "ms")
	fmt.Println("with    pool: ", endTime2-startTime2, "ms")
}

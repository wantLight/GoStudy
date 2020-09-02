package main

/**
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
*/
import "unsafe"

const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa)
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	xa = iota // 这里初始化为3
	xb
)

func main() {
	println(aa, bb, cc)
	println(xa, xb)
}

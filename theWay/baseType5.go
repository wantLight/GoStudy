package main

import (
	"fmt"
	"strings"
)

/**
这些类型的长度都是根据运行程序所在的操作系统类型所决定的：

int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
uintptr 的长度被设定为足够存放一个指针即可。
Go 语言中没有 float 类型。（Go语言中只有 float32 和 float64）没有double类型。
与操作系统架构无关的类型都有固定的大小，并在类型的名称中就可以看出来：

整数：
int8（-128 -> 127）
int16（-32768 -> 32767）
int32（-2,147,483,648 -> 2,147,483,647）
int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
无符号整数：
uint8（0 -> 255）
uint16（0 -> 65,535）
uint32（0 -> 4,294,967,295）
uint64（0 -> 18,446,744,073,709,551,615）
浮点型（IEEE-754 标准）：
float32（+- 1e-45 -> +- 3.4 * 1e38）
float64（+- 5 * 1e-324 -> 107 * 1e308）


*/
type ByteSize float64

const (
	_           = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}

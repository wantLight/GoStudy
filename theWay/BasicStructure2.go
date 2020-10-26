package main

/**
变量（或常量）包含数据，这些数据可以有不同的数据类型，简称类型。
使用 var 声明的变量的值会自动初始化为该类型的零值。类型定义了某个变量的值的集合与可对其进行操作的集合。

类型可以是基本类型，如：int、float、bool、string；
结构化的（复合的），如：struct、array、slice、map、channel；
只描述类型的行为的，如：interface。
*/
import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	fmt.Print(v)
}

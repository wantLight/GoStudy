package main

import "fmt"

/**
常量使用关键字 const 定义，用于存储不会改变的数据。
常量的值必须是能够在编译时就能够确定的
存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
常量的定义格式：const identifier [type] = value
const 关键字，iota 就重置为 0
*/

// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
const (
	a = iota // a = 0
	b        // b = 1
	d = 5    // d = 5
	e        // e = 5
)

const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

func main() {
	fmt.Println(b)
}

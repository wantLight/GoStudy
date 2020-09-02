package main

import "fmt"

/* 声明全局变量 */
var aint int = 20

// func：函数由 func 开始声明
// func function_name( [parameter list] ) [return_types]
func swap(x, y string) (string, string) {
	return y, x
}

// Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// Go 函数可以返回多个值 Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	a, b := swap("Google", "FaceBook")
	fmt.Println(a, b)

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	/* main 函数中声明局部变量 */
	// Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
	var aint int = 10
	var baint int = 20
	var c int = 0

	fmt.Printf("main()函数中 a = %d\n", aint)
	c = sum(aint, baint)
	fmt.Printf("main()函数中 c = %d\n", c)
}

/* 函数定义-两数相加 */
func sum(aint, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", aint)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return aint + b
}

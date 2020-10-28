package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	// 使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
	// a := "O"
	a = "O"
	print(a)
}

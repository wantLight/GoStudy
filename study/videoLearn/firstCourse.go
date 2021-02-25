package main

/**
1. go 语言根据函数第一个字母是否大写为public，小写private
可见性规则
当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，
如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）。

Go基本类型

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
	sm := make([]map[int]string, 5, 10)
	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "ok"
		fmt.Print(sm[i])
	}
	fmt.Println("------------")
	fmt.Println(sm)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	fmt.Println(v)

	var a = 11212
	c := int(a)
	fmt.Println(c)

}

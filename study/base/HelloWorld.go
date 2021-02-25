/**
  包声明
  引入包
  函数
  变量
  语句 & 表达式
  注释
   **/
// 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

//  import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
import (
	"fmt"
	"time"
)

//  func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
func main() {

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	//fmt.Println("Hello, World!")
	//
	//var c complex64 = 5 + 5i
	////output: (5+5i)
	//fmt.Printf("Value is: %v", c)
}

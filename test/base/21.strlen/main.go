package main

import (
	"fmt"
	"unicode/utf8"
)

// 参考：https://studygolang.com/topics/4098
// 参考：https://blog.csdn.net/sinat_32336967/article/details/95971312

func main() {

	fmt.Println(len("1二3"))
	fmt.Println(len([]byte("1二3")))
	fmt.Println(len([]rune("1二3")))
	fmt.Println(utf8.RuneCountInString("1二3"))
}

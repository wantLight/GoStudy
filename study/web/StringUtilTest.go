package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := make([]byte, 0, 100)
	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	st := string(str)
	fmt.Println(st)
	// 字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains(st, "foo"))
	// 字符串链接，把slice a通过sep链接起来
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	// 在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index(st, "abc"))
	// 重复s字符串count次，最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))
}

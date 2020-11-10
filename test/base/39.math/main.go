package main

import (
	"fmt"
	"math"
)

// 描述：四舍五入取整。
// 注意：math.Round(x)是go1.10引入的，老版本的go可以用这个方法。
// 样例：Round(4.5) = 5
// 样例：Round(4.4) = 4
// 参考：https://studygolang.com/articles/12965?fr=sidebar
func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}

func main() {
	x := 1.5
	fmt.Println(math.Ceil(x))  // 2 向上取整。
	fmt.Println(math.Floor(x)) // 1 向下取整。
	fmt.Println(math.Round(x)) // 2 四舍五入。
	fmt.Println(Round(x))      // 2 四舍五入。
}

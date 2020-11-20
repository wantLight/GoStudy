package main

import (
	"context"
	"fmt"
)

/**
如果切片的容量小于1024个元素，那么扩容的时候slice的cap就翻番，乘以2；一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。
如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果扩容之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
*/
func AddElement(slice []int, e int) []int {
	return append(slice, e)
}
func main() {
	var slice []int
	fmt.Println("容量：", cap(slice), " 长度", len(slice))
	slice = append(slice, 1, 2, 3)
	fmt.Println("容量：", cap(slice), " 长度", len(slice))
	newSlice := AddElement(slice, 4)
	fmt.Println("容量：", cap(newSlice), " 长度", len(newSlice))
	fmt.Println(&slice[0] == &newSlice[0])

	context.Background()
}

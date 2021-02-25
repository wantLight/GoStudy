package main

import (
	"fmt"
	"math"
)

// 摘自：Go 语言简介（上）— 语法 http://coolshell.cn/articles/8460.html

// ////////////////////////////////////////////////////////////////////
// 接口。
// ////////////////////////////////////////////////////////////////////
type Shape interface {
	Area() float64      // 计算面积。
	Perimeter() float64 // 计算周长。
}

// 长方形。
type Rect struct {
	width, height float64
}

// 长方形的面积。
func (r *Rect) Area() float64 {
	return r.width * r.height
}

// 长方形的周长。
func (r *Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// 圆形。
type Circle struct {
	radius float64
}

// 圆形的面积。
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// 圆形的周长
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 接口的使用
func testInterface() {

	r := Rect{width: 2.9, height: 4.8}
	c := Circle{4.3}

	shapes := []Shape{&r, &c} // 通过指针实现

	for _, shape := range shapes {
		fmt.Println(shape)
		fmt.Println(shape.Area())
		fmt.Println(shape.Perimeter())
	}
}

func main() {
	testInterface()
}

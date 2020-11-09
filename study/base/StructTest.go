package main

import "fmt"

// 摘自：Go 语言简介（上）— 语法 http://coolshell.cn/articles/8460.html

// ////////////////////////////////////////////////////////////////////
// 结构体。
// ////////////////////////////////////////////////////////////////////
type Person struct {
	Name  string
	Age   int
	Email string
}

func testStruct() {

	// 声明指针并使用"零值"进行初始化。
	pPerson1 := new(Person)
	pPerson2 := Person{}
	fmt.Println(pPerson1)
	fmt.Println(pPerson2)

	// 声明指针并使用指定的值进行初始化。
	pPerson3 := Person{"Tom", 30, "tom@gmail.com"}
	pPerson4 := Person{Name: "Tom", Age: 30, Email: "tom@gmail.com"}
	fmt.Println(pPerson3)
	fmt.Println(pPerson4)

	// 声明结构体。
	person := Person{Name: "Tom", Age: 30, Email: "tom@gmail.com"}
	fmt.Println(person) // 输出：{Tom 30 tom@gmail.com}

	// 只有 &person 才会修改原先的值
	pPerson := &person
	fmt.Println(pPerson) // 输出：&{Tom 30 tom@gmail.com}

	// go语言比较方便的一个地方是，结构体变量和结构体指针的用法相同，都是用点号(.)访问属性。
	pPerson.Age = 40
	person.Name = "Jerry"
	fmt.Println(person) // 输出：{Jerry 40 tom@gmail.com}
}

// ////////////////////////////////////////////////////////////////////
// 结构体方法。
// ////////////////////////////////////////////////////////////////////
type Rect struct {
	width, height int
}

// 求面积。
func (r *Rect) Area() int {
	return r.width * r.height
}

// 求周长。
func (r *Rect) Perimeter() int {
	return 2 * (r.width + r.height)
}

func testStructMethod() {
	r := Rect{width: 10, height: 15}

	fmt.Println("面积: ", r.Area())
	fmt.Println("周长: ", r.Perimeter())

	rp := &r
	fmt.Println("面积: ", rp.Area())
	fmt.Println("周长: ", rp.Perimeter())
}

func main() {
	testStruct()
	//testStructMethod()
}

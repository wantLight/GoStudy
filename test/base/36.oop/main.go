package main

import (
	"fmt"
)

// Go中的面向对象。摘自：https://zhuanlan.zhihu.com/p/55255987

// 封装。开始。

type User struct {
	Name string
	age  int
}

func (user *User) SetAge(age int) *User {
	user.age = age
	return user
}

func (user *User) GetAge() int {
	return user.age
}

// 封装。结束。

// 继承。开始。

type Car struct {
	Engine string
	Tire   string
}

type SUV struct {
	Car
	Brand string
}

// 继承。结束。

// 多态。开始。

type Animal interface {
	Live()
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (cat *Cat) Live() {
	fmt.Println(fmt.Sprintf("%s eat fish", cat.Name))
}

func (dog *Dog) Live() {
	fmt.Println(fmt.Sprintf("%s eat bone", dog.Name))
}

func showLive(animals []Animal) {
	for _, animal := range animals {
		animal.Live()
	}
}

// 多态。结束。

func main() {

	// 封装。
	user := User{Name: "jack"}
	user.SetAge(20)
	fmt.Println(user)

	// 继承。
	suv := SUV{Car: Car{Engine: "good", Tire: "fine"}, Brand: "BMW"}
	fmt.Println(suv.Engine)
	fmt.Println(suv.Brand)

	// 多态。
	animals := []Animal{&Cat{Name: "cat"}, &Dog{Name: "dog"}}
	showLive(animals)
}

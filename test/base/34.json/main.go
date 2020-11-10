package main

import (
	"encoding/json"
	"os"
	"runtime/debug"

	"github.com/davecgh/go-spew/spew"
)

type Student struct {
	ID   string `json:"id"   xml:"id"`
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age"  xml:"age"`
}

func main() {

	// 对象序列化。
	student1 := Student{ID: "123", Name: "张三", Age: 28}
	student1Bytes, err := json.Marshal(student1)
	if err != nil {
		spew.Printf("Json marshal student1 fail. student1=[%+v], err=[%v].\n", student1, err)
		debug.PrintStack()
		os.Exit(1)
	}
	spew.Printf("student1=[%v].\n", string(student1Bytes))

	// 对象反序列化。
	var student2 Student
	// 或使用：student2 := new(Student)
	// 或使用：student2 := &Student{}
	err = json.Unmarshal(student1Bytes, &student2)
	if err != nil {
		spew.Printf("Json unmarshal student1 fail. student1=[%v], err=[%v].\n", string(student1Bytes), err)
		debug.PrintStack()
		os.Exit(1)
	}
	spew.Printf("student2=[%+v].\n", student2)

	// 切片序列化。
	students1 := []*Student{{ID: "123", Name: "张三", Age: 28}, {ID: "456", Name: "李四", Age: 29}}
	students1Bytes, err := json.Marshal(students1)
	if err != nil {
		spew.Printf("Json marshal students1 fail. students1=[%+v], err=[%v].\n", students1, err)
		debug.PrintStack()
		os.Exit(1)
	}
	spew.Printf("students1=[%v].\n", string(students1Bytes))

	// 切片反序列化。
	var students2 []*Student
	// 或使用：students2 := new([]*Student)
	// 或使用：students2 := &[]*Student{}
	// 或使用：students2 := make([]*Student, 0)
	err = json.Unmarshal(students1Bytes, &students2)
	if err != nil {
		spew.Printf("Json unmarshal students1 fail. students1=[%v], err=[%v].\n", string(students1Bytes), err)
		debug.PrintStack()
		os.Exit(1)
	}
	spew.Printf("students2=[%+v].\n", students2)
}

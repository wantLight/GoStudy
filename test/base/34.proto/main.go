package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

/**
protobuf也叫protocol buffer是google 的一种数据交换的格式，它独立于语言，独立于平台。google 提供了多种语言的实现：java、c#、c++、go 和 python，
每一种实现都包含了相应语言的编译器以及库文件。

由于它是一种二进制的格式，比使用 xml 、json进行数据交换快许多。可以把它用于分布式应用之间的数据通信或者异构环境下的数据交换。
作为一种效率和兼容性都很优秀的二进制数据传输格式，可以用于诸如网络传输、配置文件、数据存储等诸多领域。
*/

func main() {
	s1 := &Student{} //第一个学生信息
	s1.Name = "jz01"
	s1.Age = 23
	s1.Address = "cq"
	s1.Cn = ClassName_class2 //枚举类型赋值
	ss := &Students{}
	ss.Person = append(ss.Person, s1) //将第一个学生信息添加到Students对应的切片中
	s2 := &Student{}                  //第二个学生信息
	s2.Name = "jz02"
	s2.Age = 25
	s2.Address = "cd"
	s2.Cn = ClassName_class3
	ss.Person = append(ss.Person, s2) //将第二个学生信息添加到Students对应的切片中
	ss.School = "cqu"
	fmt.Println("Students信息为：", ss)

	// Marshal takes a protocol buffer message
	// and encodes it into the wire format, returning the data.
	buffer, _ := proto.Marshal(ss)
	fmt.Println("序列化之后的信息为：", buffer)
	// 	Use UnmarshalMerge to preserve and append to existing data.
	data := &Students{}
	proto.Unmarshal(buffer, data)
	fmt.Println("反序列化之后的信息为：", data)
}

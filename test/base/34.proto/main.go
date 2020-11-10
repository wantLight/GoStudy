package main

import (
	"log"

	"github.com/xdc0209/go-base/base/34.proto/proto"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
)

// 中文乱码问题：
// 参考：https://blog.csdn.net/a_ran/article/details/80751534
// 在使用spew.Println()打印pb结构体时，如果其中的字符串包含中文的话，则输出的是八进制utf-8编码。出现这个问题的原因是pb结构体的String()方法内部使用proto.CompactTextString()实现的。
// 如果出现此种情况，可以使用此语句查看真实的字符串：fmt.Println("\351\243\216\346\270\205\346\211\254")

// 参考：https://blog.csdn.net/u014029783/article/details/80037654

func main() {

	book := &addressbook.AddressBook{}
	person := addressbook.Person{
		Id:    1234,
		Name:  "风清扬",
		Email: "mike@example.com",
		Phones: []*addressbook.Person_PhoneNumber{
			{Number: "1234-2332", Type: addressbook.Person_HOME},
		},
	}
	book.People = append(book.People, &person)
	spew.Println(book)

	bookBytes, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book.")
	}

	book2 := &addressbook.AddressBook{}
	err = proto.Unmarshal(bookBytes, book2)
	if err != nil {
		log.Fatalln("Failed to decode address book bytes.")
	}
	spew.Println(book2)
}

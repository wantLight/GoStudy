package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"sort"
)

// 参考：http://c.biancheng.net/view/81.html

type Person struct {
	Name string
	Age  int
}

// 人物排序器：优先按年龄排序，如果年龄相同，则按姓名排序。
type PersonSorter []*Person

func (s PersonSorter) Len() int {
	return len(s)
}
func (s PersonSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s PersonSorter) Less(i, j int) bool {

	if s[i].Age != s[j].Age {
		return s[i].Age < s[j].Age
	}

	if s[i].Name != s[j].Name {
		return s[i].Name < s[j].Name
	}

	return false
}

func main() {

	fmt.Println("原始数组。")
	persons := []*Person{
		{"Alice", 20},
		{"Jane", 15},
		{"Bob", 15},
	}
	// go-spew可以帮助Golang开发者打印数据的结构
	spew.Dump(persons)
	fmt.Println()

	fmt.Println("使用Sort方法升序排序。")
	sort.Sort(PersonSorter(persons))
	spew.Dump(persons)
	fmt.Println()

	fmt.Println("使用Sort方法降序排序。go中的排序默认是升序的，如果想降序，要使用sort.Reverse。")
	sort.Sort(sort.Reverse(PersonSorter(persons)))
	spew.Dump(persons)
	fmt.Println()

	fmt.Println("使用Sort方法升序、稳定排序。sort.Sort是不保证稳定排序的，如果想稳定排序，要使用sort.Stable。")
	sort.Stable(PersonSorter(persons))
	spew.Dump(persons)
	fmt.Println()

	fmt.Println("使用Sort方法降序、稳定排序。")
	sort.Stable(sort.Reverse(PersonSorter(persons)))
	spew.Dump(persons)
	fmt.Println()

	fmt.Println("使用Slice方法排序。")
	sort.Slice(persons, func(i, j int) bool {
		if persons[i].Age != persons[j].Age {
			return persons[i].Age < persons[j].Age
		}

		if persons[i].Name != persons[j].Name {
			return persons[i].Name < persons[j].Name
		}

		return false
	})
	spew.Dump(persons)
	fmt.Println()
}

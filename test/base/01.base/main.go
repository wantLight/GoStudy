package main

import (
	"fmt"
	"sort"
)

// 参考自：Go 语言简介（上）— 语法 http://coolshell.cn/articles/8460.html

// ////////////////////////////////////////////////////////////////////
// fmt输出格式。
// ////////////////////////////////////////////////////////////////////
func testPrint() {

	// go中的fmt.Printf与C中的printf用法大体相同，只是做了一些改进。

	// 基本数据类型的打印：
	// %T 打印变量的类型。
	// %v 根据变量的类型，以默认的方式打印变量的值。
	// %+v 对于基本类型(非结构体类型)，与%v相同。
	fmt.Printf("%T: %+v\n", true, true)                     // 这里的%v相当于%t。
	fmt.Printf("%T: %+v\n", 123, 123)                       // 这里的%v相当于%d。
	fmt.Printf("%T: %+v\n", 1.23, 1.23)                     // 这里的%v相当于%f。
	fmt.Printf("%T: %+v\n", "hello world.", "hello world.") // 这里的%v相当于%s。

	// 结构体数据类型的打印：
	type Person struct {
		Name  string
		Age   int
		Phone string
	}
	// %v 以默认的方式打印。比如：{sam {12345 67890}}
	// %+v 带字段名称打印。比如：{name:sam phone:{mobile:12345 office:67890}
	// %#v 用Go的语法打印。main.People{name:”sam”, phone:main.Phone{mobile:”12345”, office:”67890”}}
	// 综上，不管是基本类型，还是结构体类型，都建议使用%+v，统一处理，便于记忆，不会出错。
	fmt.Printf("%v\n", Person{"张三", 28, "13812345678"})
	fmt.Printf("%+v\n", Person{"张三", 28, "13812345678"})
	fmt.Printf("%#v\n", Person{"张三", 28, "13812345678"})

	// 字符串转义：
	// %q 字符串的首尾带有双引号，字符串中的特殊字符(", \, \n, \t等)将被转义。
	fmt.Printf("%q\n", "\\\"\n")
}

// ////////////////////////////////////////////////////////////////////
// 变量和常量。
// ////////////////////////////////////////////////////////////////////
func testVar() {

	// 变量的声明很像javascript，使用var关键字。注意：go是静态类型的语言。

	// 声明并初始化一个变量。
	var i int = 100
	var f float64 = 3.14

	// 不用指明类型，通过初始化值来推导。
	var b = true           // bool
	var s = "hello world." // string

	// 还有一种声明变量的方式。
	j := 100 // 等价于 var j int = 100

	// 声明并初始化多个变量。
	var x, y, z = 1, "two", 3

	// 常量很简单，使用const关键字：
	const pi float64 = 3.1415926

	fmt.Println(i, f, b, s, j, x, y, z, pi)
}

// ////////////////////////////////////////////////////////////////////
// 数组，数组的长度不可改变。
// ////////////////////////////////////////////////////////////////////
func testArray() {

	// 声明。注意：声明数组时必须指定长度，如不指定，则代表声明切片。
	var a [5]int
	fmt.Println("array a:", a)

	// 赋值。
	a[0] = 5
	a[1] = 10
	a[3] = 30
	fmt.Println("assign a:", a)

	// 长度。
	fmt.Println("len a:", len(a))

	// 声明并初始化。注意：声明数组时必须指定长度，如不指定，则代表声明切片。
	b := [5]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5} // 使用...代替数字，但是不能省略，否则就变成切片了。
	fmt.Printf("%T: %v\n", b, b) // 注意：b的类型为[5]int，其中数字也是类型的一部分。
	fmt.Printf("%T: %v\n", c, c) // 注意：c的类型为[5]int，其中数字也是类型的一部分。

	// 二维数组。
	var d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i + j
		}
	}
	fmt.Println("2d array d:", d)
}

// ////////////////////////////////////////////////////////////////////
// 切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。切片不能删除元素，需要删除时要创建新的切片，并拷贝其他元素。切片比数组更常用。
// ////////////////////////////////////////////////////////////////////
func testSlice() {

	var s1 []int                                                                                    // 声明切片。一个切片在未初始化之前默认为nil，长度为0。
	s2 := make([]int, 5)                                                                            // 声明并初始化切片。使用make([]T, length)函数来创建切片。
	s3 := make([]int, 5, 10)                                                                        // 声明并初始化切片。使用make([]T, length, capacity)函数来创建切片。
	fmt.Printf("s1: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s1, len(s1), cap(s1), s1) // 注意：s1的类型为[]int，中括号中是没有数字的，这是区分数组和切片的依据。
	fmt.Printf("s2: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s2, len(s2), cap(s2), s2) // 注意：s2的类型为[]int，中括号中是没有数字的，这是区分数组和切片的依据。
	fmt.Printf("s3: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s3, len(s3), cap(s3), s3) // 注意：s3的类型为[]int，中括号中是没有数字的，这是区分数组和切片的依据。

	// 切片也可以基于现有的切片或数组生成。切分的范围由两个由冒号分割的索引对应的半开区间指定。
	s4 := s3[2:4] // a[2] 和 a[3]，但不包括a[4]。
	s5 := s3[:4]  // 从 a[0]到a[4]，但不包括a[4]。
	s6 := s3[2:]  // 从 a[2]到a[4]，且包括a[2]。
	fmt.Printf("s4: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s4, len(s4), cap(s4), s4)
	fmt.Printf("s5: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s5, len(s5), cap(s5), s5)
	fmt.Printf("s6: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s6, len(s6), cap(s6), s6)

	// 追加元素。
	s1 = append(s1, 100)
	s2 = append(s2, 100)
	s3 = append(s3, 100)
	fmt.Printf("s1: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s2, len(s2), cap(s2), s2)
	fmt.Printf("s3: type=[%T], length=[%v], capacity=[%v], value=[%v]\n", s3, len(s3), cap(s3), s3)
}

// ////////////////////////////////////////////////////////////////////
// 分支循环语句。
// ////////////////////////////////////////////////////////////////////
func testBranchAndLoop() {

	// if语句。注意：if语句没有圆括号，而必需要有花括号。
	num := -1
	if num < 0 {
		// 负数
	} else if num == 0 {
		// 零
	} else {
		// 正数
	}

	// switch语句。注意：switch语句没有break，但可以使用逗号case多个值。
	num2 := 5
	switch num2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6:
		fmt.Println("four, five, six")
	default:
		fmt.Println("invalid value!")
	}

	// for语句。前面你已见过了，下面再来看看for的几种形式：（注意：Go语言中没有while）
	// 经典的for语句：init; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 精简的for语句：condition
	j := 1
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 死循环的for语句，相当于for(;;)
	k := 1
	for {
		if k > 10 {
			break
		}
		k++
	}

	// 遍历切片。
	nums := []int{1, 2, 4, 8}
	for idx, num := range nums {
		fmt.Println(idx, num)
	}
}

// ////////////////////////////////////////////////////////////////////
// map的使用，map在别的语言里可能叫哈希表或叫dict，下面是和map的相关操作的代码，代码很容易懂。
// ////////////////////////////////////////////////////////////////////
func testMap() {

	var m1 map[string]int                    // 声明map。声明之后必须初始化才能使用，向未初始化的map赋值会引起【panic: assign to entry in nil map.】，但是从未初始化的map取值不会报错，结果是取到value类型的零值。
	m2 := make(map[string]int)               // 声明并初始化切片map。使用make(map[K]V)函数来创建map。
	m3 := make(map[string]int, 10)           // 声明并初始化切片map。使用make(map[K]V, capacity)函数来创建切片。
	m4 := map[string]int{}                   // 声明并初始化切片map。
	m5 := map[string]int{"one": 1, "two": 2} // 声明并初始化切片map。
	fmt.Printf("m1: type=[%T], length=[%v], value=[%v]\n", m1, len(m1), m1)
	fmt.Printf("m2: type=[%T], length=[%v], value=[%v]\n", m2, len(m2), m2)
	fmt.Printf("m3: type=[%T], length=[%v], value=[%v]\n", m3, len(m3), m3)
	fmt.Printf("m4: type=[%T], length=[%v], value=[%v]\n", m4, len(m4), m4)
	fmt.Printf("m5: type=[%T], length=[%v], value=[%v]\n", m5, len(m5), m5)

	// 添加键值对。
	m5["three"] = 3
	fmt.Printf("%T: %v\n", m5, m5) // 输出：map[string]int: map[three:3 one:1 two:2] (顺序在运行时可能不一样)
	fmt.Println(len(m5))           // 输出：3

	// 从map里取值。
	v := m5["two"]                       // 如果key不存在返回value类型的零值。
	v, exist := m5["two"]                // 如果key不存在返回value类型的零值，exist为false。
	fmt.Printf("%T: %v\n", v, v)         // 输出：int: 2
	fmt.Printf("%T: %v\n", exist, exist) // 输出：bool: true

	// 判断key是否存在。
	if _, exist := m5["two"]; exist {
		fmt.Printf("%T: %v\n", exist, exist) // 输出：bool: true
	}

	// 删除键值对。
	delete(m5, "two")
	fmt.Println(m5) // 输出：map[three:3 one:1]

	// 遍历map(顺序不确定)。
	for key, val := range m5 { // 对于map，range返回数据的key和val。
		fmt.Printf("%s => %d\n", key, val)
	}

	// 有序遍历map。
	var keys []string
	for key, _ := range m5 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s => %d\n", key, m5[key])
	}
}

// ////////////////////////////////////////////////////////////////////
// 指针。
// ////////////////////////////////////////////////////////////////////
func testPointer() {

	var i int = 1
	var pInt *int = &i

	// 输出：i=1, pInt=0xc082052370, *pInt=1
	fmt.Printf("i=%d, pInt=%p, *pInt=%d\n", i, pInt, *pInt)

	*pInt = 2
	// 输出：i=2, pInt=0xc082052370, *pInt=2
	fmt.Printf("i=%d, pInt=%p, *pInt=%d\n", i, pInt, *pInt)

	i = 3
	// 输出：i=3, pInt=0xc082052370, *pInt=3
	fmt.Printf("i=%d, pInt=%p, *pInt=%d\n", i, pInt, *pInt)
}

// ////////////////////////////////////////////////////////////////////
// 内存分配，Go具有两个分配内存的机制，分别是内建的函数new和make。他们所做的事不同，所应用到的类型也不同，这可能引起混淆，但规则却很简单。
// ////////////////////////////////////////////////////////////////////
func testMemoryAllocation() {

	// new()和make()的区别?
	//
	// 看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。
	//
	// 详细的区别描述：
	// 1. new(T) 为类型T分配一片内存，初始化为0并且返回类型为*T的内存地址，适用于任意类型。对于结构体，new(T)相当于&T{}。
	// 2. make(T) 也是用于内存分配的，但是和new不同，它只适用于3种内建的引用类型：切片、map和channel，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	//
	// 精简的区别描述：
	// 1、make只能用来分配及初始化类型为slice，map，chan的数据；new可以分配任意类型的数据；
	// 2、new分配返回的是指针，即类型*T；make返回引用，即T；

	// 为切片结构分配内存；p为指向切片的指针；*p == nil；很少使用。
	// var p *[]int = new([]int)

	// 切片v现在是对一个新的有10个整数的数组的引用。
	// var v []int = make([]int, 10)

	// 不必要地使问题复杂化：
	var p *[]int = new([]int)
	fmt.Println(p) // 输出：&[]
	*p = make([]int, 10, 10)
	fmt.Println(p)       // 输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) // 输出：0

	// 习惯用法:
	v := make([]int, 10)
	fmt.Println(v) // 输出：[0 0 0 0 0 0 0 0 0 0]
}

// ////////////////////////////////////////////////////////////////////
// 函数，单个返回值。
// ////////////////////////////////////////////////////////////////////
func testFunc() {
	fmt.Println(max(4, 5))
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// ////////////////////////////////////////////////////////////////////
// 函数，多个返回值。
// ////////////////////////////////////////////////////////////////////
func testFuncMultiRet() {

	v, e := multiRet("one")
	fmt.Println(v, e) // 输出：1 true

	v, e = multiRet("four")
	fmt.Println(v, e) // 输出：0 false

	// 通常的用法(注意分号后有e)。
	if v, e = multiRet("four"); e {
		// 正常返回。
	} else {
		// 出错返回。
	}
}

func multiRet(key string) (int, bool) {

	m := map[string]int{"one": 1, "two": 2, "three": 3}

	var val int
	var err bool

	val, err = m[key]

	return val, err
}

// ////////////////////////////////////////////////////////////////////
// 函数，可变参数。
// ////////////////////////////////////////////////////////////////////
func testFuncVariableParameter() {
	sum(1, 2)
	sum(1, 2, 3)

	// 传切片。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums { // 对于切片，range返回数据的下标和值,下标这里不需要，使用下划线_丢掉。
		total += num
	}
	fmt.Println(nums, total)
}

func main() {
	testPrint()
	testVar()
	testArray()
	testSlice()
	testBranchAndLoop()
	testMap()
	testPointer()
	testMemoryAllocation()
	testFunc()
	testFuncMultiRet()
	testFuncVariableParameter()
}

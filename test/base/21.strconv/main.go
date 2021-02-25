package main

import (
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	// int <--> string
	i, err := strconv.Atoi("123456")
	is := strconv.Itoa(123456)

	// int64 <--> string
	// 第二个参数为基数(2~36)。
	// 第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，分别对应int, int8, int16, int32和int64。
	i64, err := strconv.ParseInt("123456", 10, 64)
	i64s := strconv.FormatInt(123456, 10)

	// uint64 <--> string
	u64, err := strconv.ParseUint("123456", 10, 64)
	u64s := strconv.FormatUint(123456, 10)

	// int、int64、uint、uint64 --> string 性能较差，不要使用：https://blog.csdn.net/flyfreelyit/article/details/79701577
	s := fmt.Sprint(123456)

	spew.Dump(i, is, i64, i64s, u64, u64s, s, err)
}

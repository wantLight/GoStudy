package main

import (
	"fmt"
	"sort"
)

// 摘自：https://www.yiibai.com/go/golang-sorting.html
// 摘自：https://gobyexample.com/sorting

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings. Note that sorting is
	// in-place, so it changes the given slice and doesn't
	// return a new one.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// bool
	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", sorted)

	// 逆序排序。go中的排序默认是升序的，如果想降序，要使用sort.Reverse。
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Reversed: ", ints)
}

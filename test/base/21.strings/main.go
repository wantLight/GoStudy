package main

import "strings"
import "fmt"

// 摘自：https://www.yiibai.com/go/golang-string-functions.html
// 摘自：https://gobyexample.com/string-functions

func main() {

	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("TrimPrefix:", strings.TrimPrefix("test", "te"))
	fmt.Println("TrimSuffix:", strings.TrimSuffix("test", "st"))
	fmt.Println("TrimSpace: ", strings.TrimSpace(" \t\n test \t\n "))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
	fmt.Println()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	fmt.Println("Len: ", len("hello"))
	fmt.Println("Char:", "hello"[1])
}

// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](http://blog.golang.org/strings)
// for more information.

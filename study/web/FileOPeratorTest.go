package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "test.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		fmt.Println(n)
		if 0 == n {
			break
		}
		fl.Write(buf[:n])
	}

	//userFile := "test.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		//fout.Write([]byte("Just a test!\r\n"))
	}
}

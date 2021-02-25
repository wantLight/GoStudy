package main

import "time"
import "fmt"

func main() {

	done := make(chan bool, 1)
	go func() {

		err := DoSomething()
		if err != nil {
			fmt.Printf("DoSomething err=[%+v].\n", err)
		}

		done <- true
	}()
	select {
	case <-done:
		fmt.Println("DoSomething finish in 3 secs.")
	case <-time.After(time.Second * 3):
		fmt.Println("DoSomething timeout after 3 secs.")
	}
}

func DoSomething() error {
	time.Sleep(time.Second * 2)
	return nil
}

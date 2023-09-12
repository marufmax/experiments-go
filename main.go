package main

import (
	"fmt"
	"time"
)

func main() {
	resultch := make(chan string)

	go func() {
		result := <-resultch
		fmt.Println("%v\n", result)
	}()

	resultch <- "Foo"
}

func fetchApi(n int) string {
	time.Sleep(time.Second * 2)

	return fmt.Sprintf("result %d", n)
}

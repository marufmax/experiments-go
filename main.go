package main

import (
	"fmt"
	"time"
)

func main() {
	result1 := fetchApi()
	result2 := fetchApi()
	result3 := fetchApi()

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
}

func fetchApi() string {
	time.Sleep(time.Second * 2)

	return "Got it"
}

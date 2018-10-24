package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	go func() {
		c <- 3
	}()
	fmt.Printf("get value %v from channel \n", <-c)
}

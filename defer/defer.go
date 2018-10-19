package main

import "fmt"

func main() {
	defer fmt.Println("Closed the file")
	fmt.Println("opened the file")

	defer fmt.Println("statement 1")
	defer fmt.Println("statement 2")
	defer fmt.Println("statement 3")
	defer fmt.Println("statement 4")
	fmt.Println("no defer statement ")
}

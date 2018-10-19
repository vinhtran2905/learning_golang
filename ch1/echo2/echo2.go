package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	//4 types of declaration
	// s:= ""
	//var s string
	// var s = ""
	//var s string = ""
}

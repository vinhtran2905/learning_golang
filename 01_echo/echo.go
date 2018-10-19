package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("example of echo of unix")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	m := make(map[string]*string)
	str := "teste"

	m["test"] = &str
}

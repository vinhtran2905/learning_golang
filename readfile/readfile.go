package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "./hello.txt"

	content, err := ioutil.ReadFile(filename)
	checkError(err)

	result := string(content)
	fmt.Println("content of this file is: ", result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

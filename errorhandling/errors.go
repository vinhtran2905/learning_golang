package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.tzt")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
		myError := errors.New("there does not exist that file")
		fmt.Println(myError)
	}

	attendance := map[string]bool{
		"Ann":  true,
		"Mike": true}
	attended, ok := attendance["Mike"]

	if ok {
		fmt.Println("Mike attended? ", attended)
	} else {
		fmt.Println("no info for that person")
	}
}

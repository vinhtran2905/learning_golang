package main

import (
	"fmt"
	"strings"
)

func main() {
	//explicitly typed declarations
	var anInteger int = 42
	var aString string = "this is a string"

	//implicitly Typed declarations by using :=
	//the prefer way

	anInteger2 := 42
	aString2 := "this is string in go"

	//constant  explicit typestring

	const PI1 float32 = 3.14

	//constant implicit . Prefer this way
	const PI2 = 3.14545

	fmt.Println("PI1: ", PI1)
	fmt.Println("PI2: ", PI2)
	fmt.Println(anInteger2, anInteger, aString, aString2)
	fmt.Printf("aString2: %v:%T \n", aString2, aString2)
	fmt.Printf("anInteger2: %v:%T\n", anInteger2, anInteger2)

	fmt.Println(strings.ToUpper("fyi"))
	fmt.Println(strings.Title(aString2))

	fmt.Println(strings.EqualFold("hello", "HELLO"))

}

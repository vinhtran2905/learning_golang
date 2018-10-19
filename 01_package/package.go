package main

import "fmt"

func main() {
	str1 := "the quick red fox"
	str2 := "jump over"
	str3 := "the lazy brown dog."
	stringLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String length:", stringLength)
	}
	//anum :=42
	//isTrue := true
}

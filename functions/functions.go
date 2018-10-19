package main

import (
	"fmt"
	"gopractice/stringutils"
)

// func Doanotherthing() {
// 	fmt.Println("doing another thing")
// }

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func main() {
	// Dosomething()
	// Doanotherthing()
	// sum := Addvalues(4, 5)
	// fmt.Println(sum)
	//
	// sum = AddAllvalues(2, 3, 4, 5)
	// fmt.Println(sum)

	f, l := stringutils.FullName("Vinh", "Tran")
	fmt.Printf("Full name is %v and length of fullname is %v\n", f, l)

	f1, l1 := stringutils.FullNameNakedReturn("Danny", "Sandoval")
	fmt.Printf("Full name is %v and length of fullname is %v\n", f1, l1)

	d := Dog{"Poo", 32, "woof"}
	fmt.Println(d)
	d.Speak()

	d.Sound = "arf"
	d.Speak()

}

// func Dosomething() {
// 	fmt.Println("Doing something")
// }
//
// func Addvalues(value1, value2 int) int {
// 	//func Addvalues(value1 int, value2 int) int {
// 	return value1 + value2
// }
//
// func AddAllvalues(values ...int) int {
// 	sum := 0
// 	for i := range values {
// 		sum += values[i]
// 	}
// 	return sum
// }

// func FullName(f, l string) (string, int) {
// 	full := f + " " + l
// 	length := len(full)
// 	return full, length
// }
//
// func FullNameNakedReturn(f, l string) (fullname string, length int) {
// 	fullname = f + " " + l
// 	length = len(fullname)
// 	//return fullname, length
// 	//naked return since specified the return from parameter
// 	return
// }

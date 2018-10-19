package main

import "fmt"

var i, j int = 1, 2

//outside function, cannot use short variable declaration
//i:= 4
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Println(SnakeReturn())
	fmt.Println(Swap("hello", "world"))
}

func SnakeReturn() (z, y string) {
	z = "hello"
	y = "gopher"
	return
}

func Swap(x string, y string) (string, string) {

	return y, x
}

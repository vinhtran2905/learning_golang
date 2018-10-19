package main

import "fmt"

func main() {

	var x byte = 4 //alias for uint8
	var y rune = 6 //alias for int32

	var flag bool = false
	var s string = "test string"

	var i int64 = -64
	var j uint64 = 64
	var c complex128 = (8 + 1i)

	var f float64 = 4.5

	fmt.Println(x, y, flag, s, i, j, c, f)
}

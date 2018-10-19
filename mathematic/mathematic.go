package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	anInt := 5
	aFloat := 5.6
	//fmt.Println("sum:", anInt+aFloat) //mismatch type . Cause invalid operation
	sum := float64(anInt) + aFloat
	fmt.Printf("sum: %v, Type: %T \n", sum, sum)

	i1, i2, i3 := 3, 4, 5

	f1, f2, f3 := 23.5, 65.1, 76.3

	floatSum := f1 + f2 + f3
	fmt.Println("sum int: ", i1+i2+i3)
	fmt.Println("sum float: ", floatSum)

	var b1, b2, b3, bigSum big.Float

	b1.SetFloat64(23.5)
	b2.SetFloat64(23.5)
	b3.SetFloat64(23.5)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("bug sum: %.10g:\n", &bigSum)

	circleRadius := 15.5

	circumrefence := circleRadius * math.Pi
	fmt.Printf("circumerence: %0.2f\n", circumrefence)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var x float64 = 0
	var result string

	if x := -42; x < 0 {
		result = "less than zero"
	} else if x == 0 {
		result = "equal to zero"
	} else {
		result = " greater than zero"
	}

	fmt.Println(result)

	//switch statement

	rand.Seed(time.Now().Unix())
	//dow := rand.Intn(7) + 1
	//fmt.Println("Day: ", dow)

	result1 := "test"
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result1 = "it's sunday"
	case 7:
		result1 = "it's saturday"

	default:
		result1 = "it's weekday"
	}
	fmt.Println(result1)

	x := -42
	switch {
	case x < 0:
		result = "less than zero"
		fallthrough
	case x == 0:
		result = " equal to zero"
	default:
		result = " greater than zero"
	}

	fmt.Println(result)

	//for loop statement
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	colors := []string{"red", "blue", "green"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}
endofprogram:
	fmt.Println("end of program")

}

package main

import (
	"fmt"
)

type Kitchen struct {
	numberOfPlate int
	//conflict variable raise
	numberOfLamp int
}

// en error happens when we have same level variable
//when use h.numberOfLamp go doest know which numberOfLamp from bedroom or Kitchen
//to fix, again, need explicitly the struct belongs of the variable such as h.Kitchen.numberOfLamp or h.kitchen.Bedroom.numberOfLamp
// type Bedroom struct {
//   numberOfLamp int
// }

//same behaviour for methods on structs

func (k *Kitchen) turnOnLamps() {
	fmt.Printf("turned on %v lamps in Kitchen\n", k.numberOfLamp)
}

type House struct {
	Kitchen
	numberOfRoom int
	numberOfLamp int
}

func (h *House) turnOnLamps() {
	fmt.Printf("Turned on %v lamps in House\n", h.numberOfLamp)
}

func main() {
	h := House{Kitchen{10, 4}, 3, 5}
	fmt.Printf("House has %v plates and %v rooms\n", h.numberOfPlate, h.numberOfRoom)
	//test conflict variabl number of lights. it will return the variable of main struct House inteads of Kitchen
	fmt.Printf("House has %v lamps\n", h.numberOfLamp)
	//if you want to get the conflict variable value from Kitchen you need specify explicitly the struct named
	fmt.Printf("The House's Kitchen has %v lamps\n", h.Kitchen.numberOfLamp)

	//which turnOnLamps() is called? from House
	h.turnOnLamps() //From House
}

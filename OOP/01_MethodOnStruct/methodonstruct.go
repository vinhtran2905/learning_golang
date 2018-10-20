package main

import (
	"fmt"
)

type Kitchen struct {
	nbOfLamps int
}

//In the above call to Area, the instance of Rectangle is passed as a value.
//You could also pass it by reference. In calling the function,
//there would be no difference whether the instance that
//you call it with is a pointer or a value because Go will automatically do
//the conversion for you.

// func (k Kitchen) turnOn() {
// 	fmt.Printf("Turned on all(%v) lamps in kitchen\n", k.nbOfLamps)
// }
func (k Kitchen) turnOnbyValue() {
	fmt.Printf("Turned on all(%v) lamps in kitchen by value\n", k.nbOfLamps)
}

func (k *Kitchen) turnOnbyReference() {
	fmt.Printf("Turned on all(%v) lamps in kitchen by reference\n", k.nbOfLamps)
}

//if we attach a function with different type in the same pacakge
//we need to define a new struct with anonymnomous type inside
type myInt struct {
	i int
}

func (i *myInt) increaseInt() {

}

func main() {
	k := Kitchen{3}
	k.turnOnbyValue()
	(&k).turnOnbyReference()
	k.turnOnbyReference()
	k.nbOfLamps = 4
	(k).turnOnbyValue()
	(&k).turnOnbyReference()
	k.turnOnbyReference()
}

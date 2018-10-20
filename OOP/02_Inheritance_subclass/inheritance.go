package main

import "fmt"

type Kitchen struct {
	nbOfLamps int
}

func (k *Kitchen) setnbOfLamps(lamps int) {
	k.nbOfLamps = lamps
}

func (k *Kitchen) getnbOfLamps() int {
	return k.nbOfLamps
}

func (k *Kitchen) turnOff() {
	fmt.Printf("turn off all lamps in kitchen\n")
}

type Bedroom struct {
	nbOfLamps int
}

func (b *Bedroom) setnbOfLamps(lamps int) {
	b.nbOfLamps = lamps
}

func (b *Bedroom) getnbOfLamps() int {
	return b.nbOfLamps
}

type House struct {
	Kitchen
	Bedroom
}

func main() {

	h := House{Kitchen{2}, Bedroom{3}}
	//ambiguous selector h.getnbOfLamps
	//	h.getnbOfLamps()

	h.Bedroom.setnbOfLamps(5)
	fmt.Printf(" house has %v lamps on Bedroom\n", h.Bedroom.getnbOfLamps())

	//house inheritance method turnOff from Kitchen.
	//actually this is just a shortcut to call the method h.Kitchen.turnOff()
	h.turnOff()
}

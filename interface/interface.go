package main

import "fmt"

type Animal interface {
	Speak(volum int) string
}

type Dog struct {
}

func (d Dog) Speak(vol int) string {
	var result string
	for i := 0; i < vol; i++ {
		result = result + "woof"
	}
	return result
}

type Cat struct {
}

func (c Cat) Speak(vol int) string {
	var result string
	for i := 0; i < vol; i++ {
		result = result + "Meow"
	}
	return result
}

func main() {
	poodle := Animal(Dog{})
	//poodle := Dog{}
	fmt.Println(poodle.Speak(3))

	animals := []Animal{Dog{}, Cat{}}

	for _, a := range animals {
		fmt.Println(a.Speak(3))
	}
}

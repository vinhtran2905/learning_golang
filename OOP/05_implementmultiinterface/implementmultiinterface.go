package main

import (
	"fmt"
)

type Swimmable interface {
	swim()
}

type Flyable interface {
	fly()
}

type Duck struct {
	name string
}

func (d Duck) swim() {
	fmt.Println("Duck is swiming")
}

func (d Duck) fly() {
	fmt.Println("Duck is flying")
}

//for scalable than java
// Let us assume now that as time evolved,
// some of the project requirements for our Duck changed -
// Duck able to squack . Our Duck now has to
// adhere to a new interface called Squakable which is distinct
// from any of the other interfaces it already implements.

// As you may notice, the functionality has been extended without any change
// to the core classes or core hierarchies. This implementation is much cleaner,
//  easily extensible, and can scale better with the changing needs of
//  the project's requirements such as in Java class Duck implements Swimable,Flyable, Squackable

type Squackable interface {
	squack()
}

func (d Duck) squack() {
	fmt.Println("squacking")
}

func main() {
	d := Duck{"White"}
	d.fly()
	d.swim()
	d.squack()
}

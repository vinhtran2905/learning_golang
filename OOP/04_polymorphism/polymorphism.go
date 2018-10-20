package main

import "fmt"

type Shaper interface {
	Area() int
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

type Square struct {
	side int
}

func (s Square) Area() int {
	return s.side * s.side
}

func main() {
	rectangle := Rectangle{3, 4}
	square := Square{4}
	fmt.Printf("Area of rectangle is %v\n", rectangle.Area())

	shaper := Shaper(rectangle)
	fmt.Printf("Area of Shaper is %v\n", shaper.Area())

	fmt.Printf("Area of Square is %v\n", square.Area())
	shaper = square //shaper = Shaper(square)
	fmt.Printf("Area of Shaper is %v\n", square.Area())

	//refactor above code by loop
	shaperArr := [...]Shaper{square, rectangle}
	fmt.Println("Looping through shapes for area ...")

	for _, s := range shaperArr {
		fmt.Printf("Shape detail %v\n", s)
		fmt.Printf("Area of shape is %v\n", s.Area())

	}

}

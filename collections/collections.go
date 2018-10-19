package main

import (
	"fmt"
	"sort"
)

func main() {
	var p *int
	//p = 5

	if p != nil {
		fmt.Println("value of p: ", *p)
	} else {
		fmt.Println("p is nil")
	}

	q := 6

	p = &q

	if p != nil {
		fmt.Println("value of p: ", *p)
	} else {
		fmt.Println("p is nil")
	}

	aFloat := 56.0
	pointer := &aFloat

	*pointer = *pointer / 7

	fmt.Println("value of pointer is ", *pointer)
	fmt.Println("value of aFloat is ", aFloat)

	aFloat = aFloat + 3

	fmt.Println("value of pointer is ", *pointer)
	fmt.Println("value of aFloat is ", aFloat)

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println(colors[2])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("number of color", len(colors))
	//array cannot add and move item
	// numbers = append(numbers, 10)
	// fmt.Println(numbers)

	//slice can add and move items
	var students = []string{"David", "Chris", "Danny"}

	fmt.Println(students)

	students = append(students, "Vinh")

	fmt.Println(students)

	//another way to define a sliceType
	flowers := make([]int, 5, 10)
	flowers[0] = 3
	flowers[1] = 34
	flowers[2] = 54
	flowers[3] = 1
	flowers[4] = 4
	fmt.Println(flowers)
	fmt.Println(cap(flowers))

	sort.Sort(sort.Reverse(sort.IntSlice(flowers)))
	fmt.Println(flowers)

	sort.Ints(flowers)
	fmt.Println(flowers)

	sort.Sort(sort.Reverse(sort.IntSlice(flowers)))
	fmt.Println(flowers)

	sort.Sort(sort.Reverse(sort.IntSlice(flowers)))
	fmt.Println(flowers)

	var p1 *int

	q1 := 10

	p1 = &q1

	fmt.Printf("value of pointer p is %v\n", *p1)
	//define a map and assign a value cause the panic, need using funcion make()
	// var m map[string]int
	// m["key"] = 42
	// fmt.Println(m)

	m := make(map[string]int)
	m["key"] = 42
	fmt.Printf("value of map m is %v \n and type of m is %T\n", m, m)
	//
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["CA"] = "California"
	states["TX"] = "Texas"
	fmt.Println(states)
	fmt.Printf("%T\n %v\n", states, states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "TX")
	fmt.Println(states)

	states["FL"] = "Florida"

	for k, v := range states {
		fmt.Printf("%v:%v\n", k, v)
	}

	type Dog struct {
		Breed  string
		Weight int
	}

	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Println(poodle.Breed)
	fmt.Printf("%+v\n", poodle)

}

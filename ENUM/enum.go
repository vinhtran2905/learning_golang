package main

import "fmt"

// for enums define a type of the enum that one want
// we want days be of type int where int ranges from 1 to 7
// Another reason is we want to be able to perform operations on Day
// adding 1/2..7 or subtracting 1/2..7 from the variables of type Day

type Day int

// next we define the constants of type Day
// we use iota to provide int values to each of the contant types
const (
	MONDAY Day = 1 + iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

var days = [...]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (d Day) String() string { return days[d-1] }

func main() {
	day := FRIDAY
	if day == FRIDAY {
		fmt.Println("Found a Friday")
	}
}

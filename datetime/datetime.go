package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2018, time.September, 11, 23, 0, 0, 0, time.UTC)
	fmt.Printf("time: %s\n", t)
	now := time.Now()
	fmt.Println(now)

	fmt.Println("the month is", t.Month())
	fmt.Println("the weekday is", now.Weekday())
}

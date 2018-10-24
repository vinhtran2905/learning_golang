package main

import (
	"fmt"
	"strconv"
	"time"
)

// All the channels we have defined until now defaults to a
// synchronous channel i.e. a datum put on the channel has to
// be taken off before another one is placed on it. Let’s implement
//  our cake making and packing factory now. Since channel communication
//   happens between goroutines, there are two aptly named functions
//    makeCakeAndSend and receiveCakeAndPack. Each receive the same reference
//     to a channel as a parameter so that they may communicate using it.

var i int

func makeCakeAndSend(cs chan string) {
	i = i + 1
	cakeName := "Strawberry cake " + strconv.Itoa(i)
	fmt.Printf("Making cake  %v and sending ....\n", cakeName)
	cs <- cakeName //send Strawberrycake
}

func receiveCakeAndPack(cs chan string) {
	cake := <-cs
	fmt.Printf("Receiving cake %v\n", cake)
}

func main() {
	//create a bidirectinal int channel
	// ic := make(chan int)

	//input value to  a bidirectinal int channel ic <-6
	//ic <- 6

	//get output from channel
	//	fmt.Printf("Output value from channel ic is %v\n", <-ic)
	cs := make(chan string)

	for i := 0; i < 3; i++ {
		go makeCakeAndSend(cs)
		go receiveCakeAndPack(cs)
		time.Sleep(1 * 1e9)
	}

}

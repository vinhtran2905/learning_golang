// we have extended our cake making factory to simulate the production of
// more than 1 flavor of cake - both strawberry and chocolate now!
// But the packing mechanism is the same. Also it takes more time to make a cake
// than pack one, which means that we can gain efficiency from packing different
// types of cake with the same packing machine. Since cakes come from different
// channels and the packer cannot know the exact moment when a cake is placed on
// either or many channels, it can use the select statement and wait across all of
//  them - as soon as one of the channels is receiving a cake/data, it will implement
//   its code block

package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false

	for {
		//if both channels are closed then we can stop
		if strbry_closed && choco_closed {
			return
		}
		fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if !strbry_ok {
				strbry_closed = true
				fmt.Println(" ... Strawberry channel closed!")
			} else {
				fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
			}
		case cakeName1, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				fmt.Println(" ... Chocolate channel closed!")
			} else {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName1)
			}
		}
	}
}

func main() {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)

	//two cake makers
	go makeCakeAndSend(choco_cs, "Chocolate", 3)   //make 3 chocolate cakes and send
	go makeCakeAndSend(strbry_cs, "Strawberry", 3) //make 3 strawberry cakes and send

	//one cake receiver and packer
	go receiveCakeAndPack(strbry_cs, choco_cs) //pack all cakes received on these cake channels

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(2 * 1e9)
}

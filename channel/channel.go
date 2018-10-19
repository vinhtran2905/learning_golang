//channel provide a way for two gorouting to communicate with each other and synchronzie their execution

package main

import (
	"fmt"
	"time"
)

//bidirection channel
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

//channel only allows to receive strings
//func pinger(c chan <- string)

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

//channel only allow to send a string
//func printer(c<- chan string)
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}

package main

import (
	"fmt"
	"time"
)

func pinger(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep()
	}
}

func main() {
	var canal chan string
	canal = make(chan string)

}

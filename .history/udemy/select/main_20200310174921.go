package main

import (
	"fmt"
	"time"
)

var (
	irc = make(chan string)
	sms = make(chan string)
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

fnuc eai(canal chan string){
	for {
		canal <- "eai"
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var canal chan string
	canal = make(chan string)

	go pinger(canal)
	go ponger(canal)
	go impressora(canal)

	var entrada string
	fmt.Scanln(&entrada)

}

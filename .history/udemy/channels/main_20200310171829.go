package main

func pinger(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ponger(canal chan string) {

}

func main() {
	var canal chan string
	canal = make(chan string)

}

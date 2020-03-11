package main

func ponger(canal chan string) {
	for {
		canal <- "ping"
	}
}

func main() {
	var canal chan string
	canal = make(chan string)

}

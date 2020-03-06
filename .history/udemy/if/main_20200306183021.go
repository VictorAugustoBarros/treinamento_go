package main

import "fmt"

func main() {
	meses := 6
	situacao := true
	cidade := "Piraquara"

	// < > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

}

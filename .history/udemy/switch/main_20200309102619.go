package main

import (
	"fmt"
	"runtime"
)

func main() {
	numero := 3
	fmt.Print("O numero ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("Você é da familia do Unix?")
	fmt.Print(runtime.GOOS)
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim!!")
	default:
		fmt.Println("Deixa essa pergunta para lá...")
	}
}

package main

import "fmt"

func main() {
	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 2

	fmt.Println("Qual a capacidade desse array?", len(teste))

	cantores := [2]string{"Michael Jackson", "John Lennon"}
	fmt.Printf("O que há nesse Array? \n %v \n", cantores)

	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("Qual a capacidade desse array?", len(capitais))
	for index, cidade := range capitais {
		fmt.Printf("Capital[%d] é %s", index, cidade)
	}
}

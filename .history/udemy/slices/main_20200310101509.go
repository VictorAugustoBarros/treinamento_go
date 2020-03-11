package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Tokyo"
	cidades[3] = "Singapura"
	fmt.Println(cidades, len(cidades), cap(cidades))
}

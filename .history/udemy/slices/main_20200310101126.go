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
}

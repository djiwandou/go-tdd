package main

import "fmt"

func main() {
	fmt.Print(AddNumber(1, 3))
}

func AddNumber(number_one int, number_two int) int {
	total := number_one + number_two
	return total
}
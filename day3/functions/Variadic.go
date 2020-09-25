package main

import "fmt"

/*
	Notes:
	-	Variadic functions can be called with any number of trailing arguments.
*/

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
}

func sum(ars ...int) int {
	sum := 0
	for _, num := range ars {
		sum += num
	}
	return sum
}

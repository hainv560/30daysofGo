package main

import "fmt"

/*
	for initialization; condition; increment {
		// loop body
	}
	Notes:
	- A for statement is used to execute a block of code repeatedly.
	- If you skip the init and increment statements, you get a while loop.
*/

func main() {

	// Three-components loop
	sum := 0
	for i := 0; i < 9; i++ {
		sum += i
	}
	fmt.Println(sum)

	// While loop
	for sum < 1000 {
		sum *= 2
	}
	fmt.Println(sum)

	// Infinite loop
	increment := 0
	for {
		increment++
		if increment < 10 {
			fmt.Println(increment)
		}
	}

}

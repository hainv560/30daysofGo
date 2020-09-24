package main

import (
	"fmt"
	"math/rand"
)

/*
	Notes:
 	- Programs start running in package main.
	- The main function in the package “main” will be the entry point of our executable program
*/
func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
}

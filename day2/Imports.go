package main

import "fmt"
import "math"

/*
	Notes:
	- It is an error to import a package without using it
	- It is good style to use the factored import statement.
		Ex: import (
				"fmt"
				"math"
			)
*/
func main() {
	fmt.Println(math.Pi)
}

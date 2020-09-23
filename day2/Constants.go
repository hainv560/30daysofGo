package main

import "fmt"

/*
	- The keyword const is used to declare a constant.
	- Constants in Go can be either typed or untyped.
		+ Typed constants work like immutable variables can inter-operate only with the same type
		+ Untyped constants work like literals can inter-operate with similar types.
*/
func main() {

	// Untyped constants
	const (
		A = 50
		B = 60
		C = "C"
	)
	const x, y = 80, 39

	fmt.Printf("%T\n%T\n", x, y)

	// Typed constants
	const z string = "ahihi"

	fmt.Printf("%T\n", z)
}

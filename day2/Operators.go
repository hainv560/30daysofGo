package main

import "fmt"

/*
	Types:
	- Arithmetic Operators: +,-,*,/,%
	- Comparison Operators: >,<,>=,<=,==,!=
	- Increment,Decrement Operators: ++,--
	- Assignment Operators: +=,-=,*=,/=,%=.....
	- Logic Operators: &&,||,!
	- Bitwise Operators: &,|,^,<<,>>
	Notes:
	- No pointer arithmetic
	- (++,--) is statement, not a expression: var x = y++ // error
	- Postfix, not prefix: ++y // Error
	- Go doesn't have a ternary operator(?:)
*/
func main() {
	var x = 5
	var y = 7

	// Arithmetic Ops
	fmt.Println("Addition=", x+y)
	fmt.Println("Subtraction", x-y)
	fmt.Println("Division", x/y)
	fmt.Println("Modulus", x%y)
	fmt.Println("Multiplication", x*y)

	// Increment,Decrement Ops
	x++
	fmt.Println("Increment", x)
	y--
	fmt.Println("Decrement", y)

	// Assignment Ops
	k := 0
	k += x
	z := 0
	z -= y
	fmt.Println(k, z)

	// Comparison Ops
	if x == y {
		fmt.Println("X equals Y")
	}

	// Logical Ops
	if x > 0 && z < 0 {
		fmt.Println("Yes! We did it")
	}

}

package main

import "fmt"

/*
	Notes:
	 - Go is a statically typed programming language. Which means that the type of every variable is known or inferred by the compiler at compile time.
	 - Using var keyword , var declare 1 or more variables at once
	 - "Inside" a function := is shorthand for declaring and initializing a variable
	 - It is an error to declare a variable without using it.
*/

func main() {

	// Declaring a single variable
	var x string
	// Assignment
	x = "ahihi"
	fmt.Println(x)

	// Declaring a variable with an initial value ,Go will automatically infer the type of that variable from the initial value.
	var a = 232
	fmt.Println(a)

	// Multiple variable declaration
	var k, m string
	fmt.Println(k, m)

	// Multiple variable declaration with initial values
	var b, c = 2, 5
	fmt.Println(b, c)

	// Multiple variable declaration with different types
	var (
		name   = "richard"
		age    = 29
		height int
	)
	fmt.Println(name, age, height)

	// Short hand declaration ,use := operator
	price, quantity := 10.3434, 34
	fmt.Println(price, quantity)

}

package main

import "fmt"

/*
	Notes:
	-	Go has support for first class functions allows functions to be assigned to variables,
		passed as arguments to other functions and returned from other functions -(Higher-order functions).
	-	Anonymous function is a function without name
*/
func main() {
	// Assigning function to the variable.
	x := func(a, b int) int {
		return a + b
	}
	fmt.Println(x(1, 2))

	// Self-Invoke
	func() {
		fmt.Println("First class function")
	}()

	// Passing arguments to anonymous functions
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 12)

	// Passing functions as arguments to other functions
	y := square(multiplication, 3)
	fmt.Println(y)

	// Returning functions from other functions
	z := transform()
	fmt.Println(z(6.2))
}

func square(callback func(a, b int) int, c int) int {
	x := callback(c, c)
	return x
}

func multiplication(a, b int) (value int) {
	value = a * b
	return
}

func convertFloatToInt(a float64) int {
	return int(a)
}

func transform() func(b float64) int {
	return convertFloatToInt
}

package main

import "fmt"

/*
	func name(parameter-list) (result-list) {
		body
	}

	Notes:
	-	Function parameters and return type(s) are optional
*/

func main() {

	// Single return value
	c := plus(1, 3)
	fmt.Println(c)

	d := plusPlus(1, 2, c)
	fmt.Println(d)

	// Multiple return values
	e, f := multipleValues(c, d)
	fmt.Println(e, f)

	// Named return values , naked return
	k, l, m := namedReturnValues(e, f, "Ahihi")
	fmt.Println(k, l, m)

	// Blank Identifier
	n, _, i := namedReturnValues(e, f, "Ahihi")
	fmt.Println(n, i)

}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multipleValues(x, y int) (int, int) {
	return x + y, 2
}

func namedReturnValues(a, b int, c string) (x1, x2 int, x3 string) {
	x1 = a + b
	x2 = a - b
	x3 = c
	return
}

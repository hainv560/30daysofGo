package main

import "fmt"

/*
	Notes:
	+	You can omit the parentheses () from an if statement in Golang, but the curly braces {} are mandatory.
	+ 	You don't need to break at the end of each case. Go's switch statements break implicitly
	+	To fall through to a subsequent case, use the fallthrough keyword
	+ 	Unlike C and Java, the case expressions do not need to be constants.
*/

func main() {

	// If with a shorthand statement
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}

	// Switch with shorthand statement
	switch day := 5; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Default")
	}

	// Multiple cases
	switch code := 4; code {
	case 1, 2, 4:
		fmt.Println("Multiple cases")
	default:
		fmt.Println("Default")

	}

	// Fall through
	switch v := 12; v {
	case 12:
		fmt.Println("12")
		fallthrough
	case 22:
		fmt.Println("22")
		fallthrough
	case 32:
		fmt.Println("32")
	default:
		fmt.Println("Default")
	}

	// Noop case: Case with no action
	switch z := 22; z {
	case 22: // Do nothing
	default:
		fmt.Println("Default")

	}

}

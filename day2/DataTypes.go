package main

import (
	"fmt"
	"reflect"
)

/*
	Types:
	- Basic types : Numbers,Booleans,Strings
	- Aggregate types: Array,Struct
	- Reference types: Pointers,Slices,Maps,Functions,Channels
	- Interface type
 	We will discuss Basic Data Types

	Notes:
	- Type conversion :  There is no automatic type promotion or conversion
		i := 55      //int
		j := 67.8    //float64
		j = i	// Error
		sum := i + j  // Error
*/

func main() {

	// Bool: true/false
	a := true
	b := false
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(b).Kind())

	/*
		Number:
		- int(Signed integers): int,int8,int16,int32,int64
		- uint(unsigned integers): uint,uint8,uint16,uint32,uint64
		- float: float32,float64
		- complex: complex64,complex128
		- byte
		- rune
	*/
	var x int8 = -89
	fmt.Println(reflect.TypeOf(x).Kind())

	var y uint32 = 43
	fmt.Println(reflect.TypeOf(y).Kind())

	z := 5.6
	fmt.Println(reflect.TypeOf(z).Kind())

	v := "This is a String"
	fmt.Println(reflect.TypeOf(v).Kind())
}

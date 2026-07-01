// package main

// import (
// 	"fmt"
// 	"math/cmplx"
// )

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) 
// 	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
// 	fmt.Printf("Type: %T Value: %v\n", z, z)
// }

 // NOTE: %T is a placeholder for the type of the variable, and %v is a placeholder for the value of the variable. The fmt.Printf function is used to format and print the output to the console.

// Types of DataTypes in Go:

// 01.Basic types: bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte (alias for uint8), rune (alias for int32), float32, float64, complex64, complex128

// 02.Aggregate types: array, struct

// 03.Reference types: pointer, function, slice, map, channel

// 04.Interface types: interface

/////////////////////////////////////////////////////////////

// ZERO VALUES

// package main

// import "fmt"

// func main() {
// 	var i int
// 	var f float64
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }

// NOTE: Go automatically initializes variables to their zero values if they are not explicitly initialized. The zero value is the default value assigned to a variable of a specific type when it is declared without an initial value.

// The zero values of the basic types in Go are as follows:
// - int: 0
// - float64: 0.0
// - bool: false
// - string: "" (empty string)

//////////////////////////////////////////////

// Type Conversions in Go

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x, y int = 3, 4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)
// 	fmt.Println(x, y, z)
// }

// NOTE: In Go, type conversions are explicit. You cannot assign a value of one type to a variable of another type without an explicit conversion. In the example above, we convert the integer values x and y to float64 before performing the square root operation, and then we convert the result back to uint before assigning it to z.


////////////////////////////////////////////////////////

//Type Inference in Go

// package main

// import "fmt"

// func main() {
// 	v := 42 // change me!
// 	fmt.Printf("v is of type %T\n", v)
// }

// NOTE: In Go, type inference allows the compiler to automatically determine the type of a variable based on the value assigned to it. In the example above, the variable v is assigned the value 42, and the compiler infers that its type is int. The %T verb in fmt.Printf is used to print the type of the variable.

///////////////////////////////////////////////////////

// Constants in Go

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// NOTE: In Go, constants are immutable values that are known at compile time and do not change during the execution of the program. They can be of any basic data type, such as boolean, numeric, or string. Constants are declared using the const keyword, and their values must be assigned at the time of declaration.

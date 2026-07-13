// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Abser interface {
// 	Abs() float64
// }

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f  // a MyFloat implements Abser
// 	a = &v // a *Vertex implements Abser

// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	a = v

// 	fmt.Println(a.Abs())
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// NOTE: An interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods. In this example, both MyFloat and *Vertex implement the Abser interface because they have the Abs() method defined. However, Vertex (without the pointer) does not implement Abser because it does not have the Abs() method defined on it directly; it only has it defined on its pointer receiver.

//////////////////////////////////////////////////////////////////

// Interfaces are implemented implicitly
// package main

// import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// // This method means type T implements the interface I,
// // but we don't need to explicitly declare that it does so.
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I = T{"hello"}
// 	i.M()
// }

// NOTE: In Go, interfaces are implemented implicitly. This means that if a type has methods that match the method signatures of an interface, it automatically implements that interface without needing to explicitly declare it. In this example, type T has a method M(), which matches the method signature of interface I, so T implements I.

////////////////////////////////////////////////////////////////

//Interface values

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t *T) M() {
// 	fmt.Println(t.S)
// }

// type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }

// func main() {
// 	var i I

// 	i = &T{"Hello"}
// 	describe(i)
// 	i.M()

// 	i = F(math.Pi)
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// NOTE: Under the hood, an interface value is represented as a tuple of a value and a concrete type. In this example, the interface I can hold values of different types (like *T and F) as long as they implement the method M(). The describe function prints the underlying value and its type, demonstrating how interface values can hold different concrete types.

/////////////////////////////////////////////////////////////////

// Interface values with nil underlying values

// package main

// import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t *T) M() {
// 	if t == nil {
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I

// 	var t *T
// 	i = t
// 	describe(i)
// 	i.M()

// 	i = &T{"hello"}
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// NOTE: If the concrete value inside an interface is nil, the interface itself is not nil. In this example, when we assign a nil pointer of type *T to the interface I, the interface holds a concrete type (*T) and a nil value. Therefore, the interface is not nil, and calling the method M() on it will still invoke the method, which checks for nil and prints "<nil>".

////////////////////////////////////////////////////////////////

// Nil interface values

// package main

// import "fmt"

// type I interface {
// 	M()
// }

// func main() {
// 	var i I
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }
 // NOTE: A nil interface value holds neither a value nor a concrete type. In this example, the variable i of type I is declared but not assigned any value, so it is nil. When we call describe(i), it prints that the value is nil and the type is <nil>. Attempting to call a method on a nil interface will result in a runtime panic because there is no concrete type to invoke the method on.

 ////////////////////////////////////////////////////////////////////

// The empty interface

// package main

// import "fmt"

// func main() {
// 	var i interface{}
// 	describe(i)

// 	i = 42
// 	describe(i)

// 	i = "hello"
// 	describe(i)
// }

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// NOTE: The empty interface, written as interface{}, can hold values of any type. In this example, the variable i is first nil, then assigned an integer value (42), and finally assigned a string value ("hello"). The describe function prints the underlying value and its type for each assignment, demonstrating the flexibility of the empty interface in Go.

/////////////////////////////////////////////////////////////////

// Type assertions

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

// NOTE: A type assertion provides access to an interface's underlying concrete value. In this example, we assert that the interface i holds a string value. The first assertion succeeds and prints "hello". The second assertion uses the "comma ok" idiom to safely check if the assertion is valid, returning true for the string and false for the float64. The last line attempts to assert a float64 from i, which causes a panic because i does not hold a float64 value.
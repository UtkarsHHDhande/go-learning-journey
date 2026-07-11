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
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

// NOTE: In Go, interfaces are implemented implicitly. This means that if a type has methods that match the method signatures of an interface, it automatically implements that interface without needing to explicitly declare it. In this example, type T has a method M(), which matches the method signature of interface I, so T implements I.

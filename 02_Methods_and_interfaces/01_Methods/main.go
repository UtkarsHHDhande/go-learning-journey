// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// }

// NOTE: Go does not have classes. However, you can define methods on types. Here, we defined a method named Abs on the Vertex type. The Abs method has a receiver of type Vertex named v. The receiver appears in its own argument list between the func keyword and the method name.

/////////////////////////////////////////////////////////

// Methods are functions

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(Abs(v))
// }

// NOTE: Remember that a method is just a function with a receiver argument. The Abs function in this example is equivalent to the Abs method in the previous example. The difference is that the Abs method has a receiver of type Vertex, while the Abs function takes a Vertex as an argument.


////////////////////////////////////////////////////////////

// Methods continued 

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func main() {
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println(f.Abs())
// }

// NOTE: You can declare a method on non-struct types as well. Here, we have declared a method named Abs on the type MyFloat. The receiver has the type MyFloat, and is named f. The Abs method returns the absolute value of f.

/////////////////////////////////////////////////////

// Pointer receivers

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(10)
// 	fmt.Println(v.Abs())
// }

// NOTE: The Scale method has a pointer receiver of type *Vertex. Since the receiver is a pointer, the Scale method can modify the value that its receiver points to. In this example, the Scale method modifies the value of v in the main function. If the Scale method had a value receiver instead, it would operate on a copy of v and would not modify the original value.
// NOTE: In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. If a method needs to modify its receiver, it must take a pointer receiver.

/////////////////////////////////////////////////////////////////

// Pointers and functions

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func Scale(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	Scale(&v, 10)
// 	fmt.Println(Abs(v))
// }

// NOTE: The Scale function takes a pointer to a Vertex as its first argument. The Scale function can modify the value that its argument points to. In this example, the Scale function modifies the value of v in the main function. If the Scale function took a Vertex as its first argument instead, it would operate on a copy of v and would not modify the original value.

///////////////////////////////////////////////////////////

//Methods and pointer indirection

// package main

// import "fmt"

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(2)
// 	ScaleFunc(&v, 10)

// 	p := &Vertex{4, 3}
// 	p.Scale(3)
// 	ScaleFunc(p, 8)

// 	fmt.Println(v, p)
// }

// NOTE: In the first call to Scale, the receiver is a value of type Vertex, and the method call is interpreted as (&v).Scale(2), taking the address of v and calling Scale with a pointer receiver. The same is true for the second call to ScaleFunc. In the third call to Scale, the receiver is already a pointer of type *Vertex, so no further indirection is necessary. The same is true for the fourth call to ScaleFunc.

/////////////////////////////////////////////////////////////

// Methods and pointer indirection (2)

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func AbsFunc(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// 	fmt.Println(AbsFunc(v))

// 	p := &Vertex{4, 3}
// 	fmt.Println(p.Abs())
// 	fmt.Println(AbsFunc(*p))
// }

// NOTE: In the first call to Abs, the receiver is a value of type Vertex, and the method call is interpreted as (&v).Abs(), taking the address of v and calling Abs with a pointer receiver. The same is true for the second call to AbsFunc. In the third call to Abs, the receiver is already a pointer of type *Vertex, so no further indirection is necessary. The same is true for the fourth call to AbsFunc.

//////////////////////////////////////////////////////////////

// Choosing a value or pointer receiver

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}


// NOTE: In this example, we have chosen to use pointer receivers for both the Scale and Abs methods. This allows us to modify the value of v in the main function when calling Scale, and also allows us to call Abs on a pointer receiver without needing to take the address of v.
// NOTE: In general, you should use pointer receivers if the method needs to modify the receiver, or if the receiver is a large struct and you want to avoid copying it. Use value receivers if the method does not need to modify the receiver and the receiver is small and inexpensive to copy.
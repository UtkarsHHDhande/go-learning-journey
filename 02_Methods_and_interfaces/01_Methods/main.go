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

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// NOTE: You can declare a method on non-struct types as well. Here, we have declared a method named Abs on the type MyFloat. The receiver has the type MyFloat, and is named f. The Abs method returns the absolute value of f.
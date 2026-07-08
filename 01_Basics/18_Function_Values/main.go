// package main

// import (
// 	"fmt"
// 	"math"
// )

// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// func main() {
// 	hypot := func(x, y float64) float64 {
// 		return math.Sqrt(x*x + y*y)
// 	}
// 	fmt.Println(hypot(5, 12))

// 	fmt.Println(compute(hypot))
// 	fmt.Println(compute(math.Pow))
// }

// NOTE : Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values. In this example, we compute the hypotenuse of a right triangle using a function literal and the math.Pow function. The compute function takes a function as an argument and applies it to the numbers 3 and 4.

///////////////////////////////////////////////////////////////////

// Function Closures

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// NOTE : A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables. In this example, adder returns a closure. Each closure is bound to its own sum variable.
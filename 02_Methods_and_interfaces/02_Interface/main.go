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

// package main

// import "fmt"

// func main() {
// 	var i interface{} = "hello"

// 	s := i.(string)
// 	fmt.Println(s)

// 	s, ok := i.(string)
// 	fmt.Println(s, ok)

// 	f, ok := i.(float64)
// 	fmt.Println(f, ok)

// 	f = i.(float64) // panic
// 	fmt.Println(f)
// }

// NOTE: A type assertion provides access to an interface's underlying concrete value. In this example, we assert that the interface i holds a string value. The first assertion succeeds and prints "hello". The second assertion uses the "comma ok" idiom to safely check if the assertion is valid, returning true for the string and false for the float64. The last line attempts to assert a float64 from i, which causes a panic because i does not hold a float64 value.

/////////////////////////////////////////////////////////////////

// Type switches

// package main

// import "fmt"

// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }

// func main() {
// 	do(21)
// 	do("hello")
// 	do(true)
// }

// NOTE: A type switch is a construct that allows you to compare the dynamic type of an interface value against multiple types. In this example, the function do takes an empty interface as an argument and uses a type switch to determine the underlying type of the value. Depending on whether the value is an int, string, or some other type, it executes different code blocks. The default case handles any types that are not explicitly matched in the switch cases.

/////////////////////////////////////////////////////////

//Stringers

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }

//NOTE: One of the most ubiquitous interfaces is Stringer defined by the fmt package.
// type Stringer interface {
//     String() string
// }
// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

///////////////////////////////////////////////////////////

// Exercise: Stringers

//Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

// package main

// import "fmt"

// type IPAddr [4]byte

// // Implement the Stringer interface
// func (ip IPAddr) String() string {
// 	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
// }

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}
// }

/////////////////////////////////////////////////////////////

// Errors

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s",
// 		e.When, e.What)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// NOTE: The error type is a built-in interface similar to fmt.Stringer. The error interface type is implemented by any value that has an Error() string method. In this example, MyError implements the error interface by defining the Error() method, allowing it to be used as an error value. The run function returns a pointer to a MyError instance, and in main, we check for an error and print it if it exists.

//////////////////////////////////////////////////////////////

//Exercise: Errors

// Copy your Sqrt function from the earlier exercise and modify it to return an error value.

// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

// Create a new type

// type ErrNegativeSqrt float64
// and make it an error by giving it a

// func (e ErrNegativeSqrt) Error() string
// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type ErrNegativeSqrt float64

// // Implement the Error() method
// func (e ErrNegativeSqrt) Error() string {
// 	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
// }

// func Sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, ErrNegativeSqrt(x)
// 	}
// 	return math.Sqrt(x), nil
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(Sqrt(-2))
// }

////////////////////////////////////////////////////////////////////

//Readers

// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// NOTE: The io package specifies the io.Reader interface, which wraps the basic Read method. The strings.NewReader function returns a new Reader reading from the provided string. In this example, we read from the string in chunks of 8 bytes, printing the number of bytes read, any error encountered, and the contents of the buffer. When the end of the string is reached, io.EOF is returned to indicate that there is no more data to read.

//////////////////////////////////////////////////////////////////////

//Exercise: Readers

//Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
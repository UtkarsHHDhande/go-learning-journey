// package main

// import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// 	Z int
// }

// func main() {
// 	fmt.Println(Vertex{1, 2, 3})
// }

// NOTE: A struct is a collection of fields. Structs are useful for grouping data together to form records.

//////////////////////////////////////////////////////////

// Struct Fields

// package main

// import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	v.X = 4
// 	fmt.Println(v.X)
// }

// NOTE: The members of a struct are accessed using a dot.

///////////////////////////////////////////////////////

// Pointers to Structs

// package main

// import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.X = 1e9
// 	fmt.Println(v)
// }

// NOTE: Struct fields can be accessed through a struct pointer. The expression (*p).X is equivalent to p.X.

//////////////////////////////////////////////////////////////

// Struct Literals

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

// NOTE: A struct literal denotes a newly allocated struct value by listing the values of its fields. You can list just a subset of fields by using the Name: syntax. The special prefix "&" returns a pointer to the struct value.
// package main

// import "fmt"

// func main() {
// 	var a [2]string
// 	a[0] = "Hello"
// 	a[1] = "World"
// 	fmt.Println(a[0], a[1])
// 	fmt.Println(a)

// 	primes := [6]int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(primes)
// }

// NOTE: The type of an array is [n]T, where n is the number of elements and T is the element type. An array's length is part of its type, so arrays cannot be resized.


////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}

// 	var s []int = primes[1:4]
// 	fmt.Println(s)
// }

// NOTE: Slices are a key data type in Go, giving a more powerful interface to sequences than arrays. A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array.

/////////////////////////////////////////////////////////////////

// Slices are like references to arrays

// package main

// import "fmt"

// func main() {
// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names)

// 	a := names[0:2]
// 	b := names[1:3]
// 	fmt.Println(a, b)

// 	b[0] = "XXX"
// 	fmt.Println(a, b)
// 	fmt.Println(names)
// }

// NOTE: Slices are like references to arrays. This means that if you change the elements of a slice, you change the corresponding elements of its underlying array.

////////////////////////////////////////////////////////////////////////

// Slice literals


// package main

// import "fmt"

// func main() {
// 	q := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(q)

// 	r := []bool{true, false, true, true, false, true}
// 	fmt.Println(r)

// 	s := []struct {
// 		i int
// 		b bool
// 	}{
// 		{2, true},
// 		{3, false},
// 		{5, true},
// 		{7, true},
// 		{11, false},
// 		{13, true},
// 	}
// 	fmt.Println(s)
// }

// NOTE: Slice literals are like array literals without the length. This means that they are more flexible than array literals, as they can be resized.


////////////////////////////////////////////////////////////////////////////

// Slice defaults

// package main

// import "fmt"

// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}

// 	s = s[1:4]
// 	fmt.Println(s)

// 	s = s[:2]
// 	fmt.Println(s)

// 	s = s[1:]
// 	fmt.Println(s)
// }

// NOTE: When slicing, you may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound.

///////////////////////////////////////////////////////////////////////////////

// // Slice length and capacity

// package main

// import "fmt"

// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}
// 	printSlice(s)

// 	// Slice the slice to give it zero length.
// 	s = s[:0]
// 	printSlice(s)

// 	// Extend its length.
// 	s = s[:4]
// 	printSlice(s)

// 	// Drop its first two values.
// 	s = s[2:]
// 	printSlice(s)
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }


// NOTE: The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

/////////////////////////////////////////////////////////////////////////////////

// Nil slices

// package main

// import "fmt"

// func main() {
// 	var s []int
// 	fmt.Println(s, len(s), cap(s))
// 	if s == nil {
// 		fmt.Println("nil!")
// 	}
// }

// NOTE: A nil slice has a length and capacity of 0 and has no underlying array. The nil value is useful to distinguish between an uninitialized slice and an empty slice.


////////////////////////////////////////////////////////////////////////////

// Creating a slice with make


// package main

// import "fmt"

// func main() {
// 	a := make([]int, 5)
// 	printSlice("a", a)

// 	b := make([]int, 0, 5)
// 	printSlice("b", b)

// 	c := b[:2]
// 	printSlice("c", c)

// 	d := c[2:5]
// 	printSlice("d", d)
// }

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// NOTE: The built-in function make is used to create slices, as well as maps and channels. We can specify both a length and a capacity when creating a slice with make. If the capacity is omitted, it defaults to the specified length.

///////////////////////////////////////////////////////////////////////////////////

// Slices of slices

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Create a tic-tac-toe board.
// 	board := [][]string{
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 	}

// 	// The players take turns.
// 	board[0][0] = "X"
// 	board[2][2] = "O"
// 	board[1][2] = "X"
// 	board[1][0] = "O"
// 	board[0][2] = "X"

// 	for i := 0; i < len(board); i++ {
// 		fmt.Printf("%s\n", strings.Join(board[i], " "))
// 	}
// }

// NOTE: Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.

///////////////////////////////////////////////////////////////////////////////

// Appending to a slice

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// NOTE: Slices can be grown dynamically. We can append new elements to a slice, and it will grow as needed. If the backing array of the slice is too small to fit the new elements, a bigger array will be allocated. The returned slice will point to the newly allocated array.

////////////////////////////////////////////////////////////////////////////////////////////////////////





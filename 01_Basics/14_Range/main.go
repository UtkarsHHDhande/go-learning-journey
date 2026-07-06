// package main

// import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// func main() {
// 	for i, v := range pow {
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }


// NOTE: The range form of the for loop iterates over a slice or map. When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.


//////////////////////////////////////////////////////////////

// Range continued

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// NOTE: You can skip the index or value by assigning to _.
// for i, _ := range pow { ... } // skip value
// for _, value := range pow { ... } // skip index
// If you only want the index, drop the ", value" entirely: for i := range pow { ... }.

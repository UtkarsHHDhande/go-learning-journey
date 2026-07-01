// package main

// import "fmt"

// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// NOTE: In Go, the for loop is the only looping construct. It can be used in various forms, including the traditional for loop with initialization, condition, and post statements, as shown in the example above.


// package main

// import "fmt"

// func main() {
// 	sum := 1
// 	for ; sum < 1000; {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

// NOTE: The init and post statements are optional.



// package main

// import "fmt"

// func main() {
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

// NOTE:  For is Go's "while"
// At that point you can drop the semicolons: C's while is spelled for in Go.


package main

func main() {
	for {
	}
}

// NOTE: Forever
// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
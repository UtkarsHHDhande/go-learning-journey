// Goroutines

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world")
// 	say("hello")
// }


// NOTE: Goroutines are lightweight threads managed by the Go runtime. When you call a function with the `go` keyword, it runs concurrently in its own goroutine. In this example, `say("world")` runs in a separate goroutine, while `say("hello")` runs in the main goroutine. The program will print "hello" and "world" concurrently, but since the main function waits for `say("hello")` to finish, the program will exit after that, potentially before "world" has finished printing.

///////////////////////////////////////////////////////////////

// Channels

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

//NOTE: Channels are a way for goroutines to communicate with each other and synchronize their execution. In this example, we create a channel `c` of type `int`. The `sum` function calculates the sum of a slice of integers and sends the result to the channel. In the `main` function, we start two goroutines to calculate the sum of two halves of the slice concurrently. We then receive the results from the channel and print them.
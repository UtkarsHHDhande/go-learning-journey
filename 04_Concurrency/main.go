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

// package main

// import "fmt"

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // receive from c

// 	fmt.Println(x, y, x+y)
// }

//NOTE: Channels are a way for goroutines to communicate with each other and synchronize their execution. In this example, we create a channel `c` of type `int`. The `sum` function calculates the sum of a slice of integers and sends the result to the channel. In the `main` function, we start two goroutines to calculate the sum of two halves of the slice concurrently. We then receive the results from the channel and print them.

//////////////////////////////////////////////////////////////////

// Buffered Channels

// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

//NOTE: Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
// ch := make(chan int, 100)
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

/////////////////////////////////////////////////////////////////

//Range and Close

// package main

// import (
// 	"fmt"
// )

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

//NOTE: A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: `v, ok := <-ch`. If `ok` is false, then the channel is closed and no more values will be sent. The `range` form of the receive loop automatically terminates when the channel is closed.

///////////////////////////////////////////////////////////////

// Select

// package main

// import "fmt"

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }

// NOTE: The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. If multiple cases can proceed, then one case is chosen at random to execute. The `select` statement is like a `switch`, but for channels. In this example, the `fibonacci` function sends Fibonacci numbers to channel `c` until it receives a signal on the `quit` channel, at which point it prints "quit" and returns.

/////////////////////////////////////////////////////////////////////

// Default Selection

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// NOTE: The default case in a select is run if no other case is ready.

// Use a default case to try a send or receive without blocking:

// select {
// case i := <-c:
//     // use i
// default:
//     // receiving from c would block
// }
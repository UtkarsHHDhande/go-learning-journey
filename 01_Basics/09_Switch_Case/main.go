// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Print("Go runs on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("macOS.")
// 	case "linux":
// 		fmt.Println("Linux.")
// 	default:
// 		// freebsd, openbsd,
// 		// plan9, windows...
// 		fmt.Printf("%s.\n", os)
// 	}
// }

// A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

///////////////////////////////////////////////////////////

// switch evaluation order

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("When's Saturday?")
// 	today := time.Now().Weekday()
// 	switch time.Saturday {
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// }

// NOTE: The cases in a switch are evaluated from top to bottom, stopping when a case succeeds. In this example, if today is Friday, the first case will succeed, and the program will print "Tomorrow." If today is Thursday, the first two cases will fail, and the third case will succeed, printing "In two days."

//////////////////////////////////////////////////////////

// switch with no condition

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// NOTE: You can also use a switch statement without an expression. This is a way to express if/else logic. Here we also use a short statement to execute before the switch.

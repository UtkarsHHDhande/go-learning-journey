package main

import "fmt"

// NOTE: Functions are defined using the func keyword, followed by the function name, a list of parameters in parentheses, and the return type. The function body is enclosed in curly braces.
// func add(x int,y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(369, 108))
// }

// NOTE: If two or more consecutive parameters share the same type, you can omit the type name for the earlier parameters. The following function is equivalent to the previous one.

// func swap(x,y string) (string, string) {
// 	return y, x
// }

// NOTE: A function can return any number of results. The swap function returns two strings, which are the two input strings in reverse order.

// func main() {
// 	a,b := swap("hello", "world")
// 	fmt.Println(a,b)
// }

// NOTE: Go's return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values.
func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return x,y
}

func main() {
	fmt.Println(split(17))
}

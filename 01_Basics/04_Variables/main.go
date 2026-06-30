package main

import "fmt"

// NOTE: Variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls. A var statement declares a list of variables; as in function argument lists, the type is last. A var statement can be at package or function level. We see that the variable i is declared with the type int and initialized to 7.

// var c, python, java bool

// func main() {
// 	i := 7
// 	fmt.Println(i, c, python, java)
// }


//NOTE: Variables can be initialized at the time of declaration. If an initializer is present, the type can be omitted; the variable will take the type of the initializer. The := syntax is shorthand for declaring and initializing a variable, e.g. var i int = 7 can be written as i := 7.


// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }


// NOTE: The := syntax is shorthand for declaring and initializing a variable, e.g. var i int = 7 can be written as i := 7. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

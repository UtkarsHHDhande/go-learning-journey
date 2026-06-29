package main

//NOTE: package name should be small letters, and it should match the directory name. The main package is a special package that defines a standalone executable program, rather than a shared library.

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(50))
	fmt.Println(math.Pi)
	//NOTE: package data types are exported if they start with a capital letter. For example, math.Pi is exported, while math.pi is not.
}

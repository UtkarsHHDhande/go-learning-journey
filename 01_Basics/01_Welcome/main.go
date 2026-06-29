package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界") // Go supports utf-8 characters
} 

// GOOS=windows GOARCH=amd64 go build -o hello.exe main.go
// this command will build the executable for Windows on an AMD64 architecture, regardless of the host operating system.
package main

import "fmt"

// Hello creates a hello message
func Hello() string {
	return "Hello World!"
}

func main() {
	msg := Hello()
	fmt.Println(msg)
}

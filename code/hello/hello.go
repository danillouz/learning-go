package main

import "fmt"

// Hello creates a hello message
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return "Hello " + name
}

func main() {
	msg := Hello("World")
	fmt.Println(msg)
}

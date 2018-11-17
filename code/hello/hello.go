package main

import "fmt"

// Hello creates a hello message
func Hello(lang, name string) string {
	var prefix string

	switch lang {
	case "English":
		prefix = "Hello "

	case "Dutch":
		prefix = "Hallo "
	}

	if name == "" {
		name = "Go"
	}

	return prefix + name
}

func main() {
	msg := Hello("English", "World")
	fmt.Println(msg)
}

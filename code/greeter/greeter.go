package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Greet prints a greeting.
func Greet(w io.Writer, greeting string) {
	fmt.Fprint(w, greeting)
}

func main() {
	Greet(os.Stdout, "it works!\n")

	handler := func(w http.ResponseWriter, r *http.Request) {
		Greet(w, "this also works!\n")
	}

	http.ListenAndServe(
		":8888",
		http.HandlerFunc(handler),
	)
}

package main // Defines the main package for an executable program.

import (
	"fmt"         // Standard library for formatted I/O.
	"rsc.io/quote" // External package to retrieve a quote.
)

func main() {
	fmt.Println("Hello world")    // Prints a basic greeting.
	fmt.Println(quote.Hello())    // Uses a function from the external package to print a quote.
}
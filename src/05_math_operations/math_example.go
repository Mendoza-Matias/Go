package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare two variables for basic arithmetic operations
	a := 10
	b := 20

	// Decrement variable 'b' by 2
	b--  // b = 19
	b--  // b = 18

	// Increment 'a' by 5
	a += 5  // a = 15

	// Print the value of 'a'
	fmt.Println("Value of a:", a)

	// Print the value of Pi from the math package
	fmt.Print("Value of Pi: ", math.Pi)

	// Print the square root of 64
	fmt.Print("Square root of 64: ", math.Sqrt(64))
}

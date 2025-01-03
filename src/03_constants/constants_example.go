// First Constants Example
package main

import (
	"fmt"
)

// Package-level constants
const Pi float32 = 3.14 // Declare and initialize a constant

const (
	x int = 100 // Single constant declaration
)

const (
	/*
		iota is used to generate a sequence of constants.
		It starts at 0 and increments by 1 automatically for each line.
	*/
	Lunes     int = iota + 1 // 0 + 1 = 1
	Martes                   // 2
	Miercoles                // 3
	Jueves                   // 4
	Viernes                  // 5
)

func main() {
	// Function-level constant
	const name string = "Hector"

	// Print constants
	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Day of the week (Martes):", Martes)
	fmt.Println("Name constant:", name)
}

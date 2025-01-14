package main

import (
	"fmt"
)

func main() {

	// Test the `split` function with various inputs
	// The third call will generate a panic due to division by zero
	split(100, 10) // Expected output: 10
	split(200, 25) // Expected output: 8
	split(34, 0)   // Will trigger a panic, handled by the `recover` mechanism
}

// `split` performs division and demonstrates panic handling
func split(dividend, divider int) {
	// Defer a function to recover from a panic if it occurs
	defer func() {
		/*
			Handle exceptional cases where a panic is triggered.
			This deferred function recovers the program flow by
			capturing the panic value, allowing the program to continue.
		*/
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Validate the input to prevent division by zero
	validateZero(divider)

	// Perform the division and print the result
	fmt.Println(dividend / divider)
}

// `validateZero` checks if the divider is zero and triggers a panic if it is
func validateZero(divider int) {
	if divider == 0 {
		// Trigger a panic when attempting to divide by zero
		panic("cannot be divided by zero")
	}
}

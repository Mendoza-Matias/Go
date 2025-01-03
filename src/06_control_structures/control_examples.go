package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// Conditional Statements

	// Example of if-else statement based on the current time
	if t := time.Now(); t.Hour() < 12 {
		// If it's before noon, print "tomorrow"
		fmt.Println("morning")
	} else if t.Hour() < 17 {
		// If it's between noon and 5 PM, print "late"
		fmt.Println("afternoon")
	} else {
		// If it's after 5 PM, print "night"
		fmt.Println("night")
	}

	// Switch statement to determine the operating system
	os := runtime.GOOS // Retrieve the operating system

	switch os {
	case "windows":
		// If OS is Windows, print this message
		fmt.Println("Running on Windows")
	case "linux":
		// If OS is Linux, print this message
		fmt.Println("Running on Linux")
	default:
		// If it's any other OS, print this message
		fmt.Println("Running on other OS")
	}

	// Switch statement with a variable initialization and case checks
	switch num := 10; num {
	case 10:
		// If num is 10, print "ten"
		fmt.Println("ten")
	case 7:
		// If num is 7, print "seven"
		fmt.Println("seven")
	default:
		// For any other number, print "other number"
		fmt.Println("other number")
	}

	// ------ FIRST LOOP ------------------------

	// For loop with condition and manual incrementation
	var i int
	for i <= 10 {
		fmt.Print(" ", i) // Print the current value of i
		i++
		if i == 6 {
			break // Exit the loop when i equals 6
		}
	}
	fmt.Println("") // Move to the next line

	// For loop with continue to skip specific values
	for i := 0; i <= 6; i++ {
		if equal2(i) {
			// Skip the iteration when i equals 2
			continue
		}
		// Print the current value of i, excluding 2
		fmt.Print(" ", i) // Output: 0 1 3 4 5 6
	}
	fmt.Println("") // Move to the next line

	// Call a function that returns two values and print them
	s, r := twoResult(5, 2)
	fmt.Println("Sum =>", s)
	fmt.Println("Difference =>", r)
}

// Function to check if a number is equal to 2
func equal2(number int) bool {
	return number == 2
}

/*
Function that returns two values (sum and difference).
Although returning multiple values is common in Go, it's often better to return
a value and an error for clearer handling in larger applications.
*/
func twoResult(n1 int, n2 int) (s, r int) {
	// Calculate sum and difference
	s = n1 + n2
	r = n1 - n2
	// Return both values
	return s, r
}

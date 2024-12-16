package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Variables with default values
	var (
		defaultInt    int     // 0
		defaultUint   uint    // 0
		defaultFloat  float32 // 0
		defaultBool   bool    // false
		defaultString string  // ""
	)

	// Print default values
	fmt.Println(defaultInt)
	fmt.Println(defaultUint)
	fmt.Println(defaultFloat)
	fmt.Println(defaultBool)
	fmt.Println(defaultString)

	// Example of type conversion between integers
	var integer16 int16 = 50
	var integer32 int32 = 100

	// Convert integer16 to int32 to operate with integer32
	fmt.Println(int32(integer16) + integer32)

	// Convert a string to an integer
	s := "100"
	i, _ := strconv.Atoi(s) // Atoi converts a string to an integer

	fmt.Println(i)

	// Convert an integer to a string
	n := 42
	s = strconv.Itoa(n) // Itoa converts an integer to a string

	// Concatenate strings
	fmt.Println(s + s)
}

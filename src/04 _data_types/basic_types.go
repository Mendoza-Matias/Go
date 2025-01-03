package main

import (
	"fmt"
	"strconv" // Package for string-to-integer and integer-to-string conversions
)

func main() {
	// Variables with default zero values
	var (
		defaultInt    int     // Defaults to 0
		defaultUint   uint    // Defaults to 0
		defaultFloat  float32 // Defaults to 0
		defaultBool   bool    // Defaults to false
		defaultString string  // Defaults to an empty string ""
	)

	// Print default values of different types
	fmt.Println("Default int:", defaultInt)
	fmt.Println("Default uint:", defaultUint)
	fmt.Println("Default float:", defaultFloat)
	fmt.Println("Default bool:", defaultBool)
	fmt.Println("Default string:", defaultString)

	// Example: Type conversion between integers
	var integer16 int16 = 50
	var integer32 int32 = 100

	// Convert `int16` to `int32` to perform operations
	fmt.Println("Sum of int16 and int32:", int32(integer16)+integer32)

	// Example: String to integer conversion
	s := "100"
	i, _ := strconv.Atoi(s) // Converts string to integer; ignores error for simplicity
	fmt.Println("Converted string to integer:", i)

	// Example: Integer to string conversion
	n := 42
	s = strconv.Itoa(n) // Converts integer to string
	fmt.Println("Converted integer to string and concatenated:", s+s) // Concatenates strings
}

package main

import "fmt"

func main() {
	// Defining a map to store colors and their hexadecimal codes
	colors := map[string]string{
		"red":   "#FF0000",
		"black": "#000000",
		"blue":  "#0000FF",
	}

	// Print the entire map
	fmt.Println("Colors map:", colors)

	// Access a specific key in the map
	fmt.Println("Hex code for 'red':", colors["red"])

	// Check if a key exists
	value, exists := colors["white"]
	if exists {
		// If the key exists, print its value
		fmt.Println("Hex code for 'white':", value)
	} else {
		// If the key does not exist, print this message
		fmt.Println("Key 'white' does not exist.")
	}

	fmt.Println("-----------------------------")

	// Inline check for key existence (for 'red')
	if value, exists := colors["red"]; exists {
		// If 'red' exists, print its value
		fmt.Println("Key 'red' exists with value:", value)
	} else {
		// If 'red' does not exist, print this message
		fmt.Println("Key 'red' does not exist.")
	}

	fmt.Println("------------------------------")

	// Delete a key-value pair from the map (removing 'blue')
	delete(colors, "blue")
	// Print the map after the deletion
	fmt.Println("After deleting 'blue':", colors)

	fmt.Println("-----------------------------")

	// Iterate over the map using a for loop
	fmt.Println("Iterating over the map:")
	for key, value := range colors {
		// Print each key-value pair
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

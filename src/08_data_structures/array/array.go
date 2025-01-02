package main

import "fmt"

func main() {
	// Declaration of 1D arrays
	var array1 [5]int                             // Array with fixed length of 5
	var array2 = [5]int{20, 30, 40, 50, 60}       // Explicit initialization
	var array3 = [...]int{100, 200, 300, 400, 50} // Length inferred automatically

	// Assigning values to array1
	array1[0] = 10
	array1[1] = 20

	// Printing arrays and their lengths
	fmt.Println("Array 1:", array1)
	fmt.Println("Length of Array 1:", len(array1))
	fmt.Println("Array 2:", array2)
	fmt.Println("Array 3:", array3)

	// Accessing a specific element
	fmt.Println("First element of Array 1:", array1[0]) // Returns 10

	// Iterating over array1 using a traditional for loop
	fmt.Print("Elements of Array 1:")
	for i := 0; i < len(array1); i++ {
		fmt.Print(" ", array1[i])
	}
	fmt.Println()

	// Iterating over array2 using a for-range loop
	fmt.Println("Elements of Array 2 with indices:")
	for index, value := range array2 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Declaration of a 2D array (matrix)
	var matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Printing the 2D array (matrix)
	fmt.Println("2D Array (Matrix):")
	for _, row := range matrix {
		fmt.Println(row)
	}
}

package main

import "fmt"

func main() {
	// Example 1: Basic slice operations
	years := []int{2002, 2003, 2004, 2005, 2006}

	// Create a new slice from an existing slice (first 4 elements)
	yearsSubset := years[:4] // Slicing from index 0 to 3 (exclusive)
	fmt.Println("Original slice (years):", years)
	fmt.Println("Subset slice (yearsSubset):", yearsSubset)

	// Append to a slice (adding 2014 to the yearsSubset)
	updatedYears := append(yearsSubset, 2014)
	fmt.Println("Updated slice after append (updatedYears):", updatedYears)

	// Print capacity and length of slices
	fmt.Println("Capacity of years:", cap(years))               // Capacity of original slice
	fmt.Println("Length of years:", len(years))                 // Length of original slice
	fmt.Println("Capacity of updatedYears:", cap(updatedYears)) // Capacity after append

	fmt.Println("--------------------------------")

	// Example 2: Remove an element from a slice
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Original days of the week:", daysOfWeek)

	// Remove "Wednesday" (index 2)
	// We combine the elements before and after the "Wednesday" element
	updatedDays := append(daysOfWeek[:2], daysOfWeek[3:]...)
	fmt.Println("Updated days of the week (without Wednesday):", updatedDays)

	fmt.Println("--------------------------------")

	// Example 3: Create slices using make
	names := make([]string, 4, 10) // Slice with length 4 and capacity 10
	names[0] = "Matias"
	names[1] = "John"
	names[2] = "Alice"
	names[3] = "Sophia"

	// Print the names slice and its length and capacity
	fmt.Println("Names slice:", names)
	fmt.Println("Length of names:", len(names))
	fmt.Println("Capacity of names:", cap(names))

	fmt.Println("--------------------------------")

	// Example 4: Copy slices
	sourceSlice := []int{1, 2, 3, 4, 5}
	targetSlice := make([]int, len(sourceSlice)) // Create a target slice with the same length as sourceSlice
	copy(targetSlice, sourceSlice)               // Copy the elements from source to target

	// Print both slices
	fmt.Println("Source slice:", sourceSlice)
	fmt.Println("Target slice after copy:", targetSlice)

	fmt.Println("--------------------------------")

	// Example 5: Append multiple elements from one slice to another
	numbers := []int{10, 20, 30}
	moreNumbers := []int{40, 50, 60}

	// Combine the two slices using the ellipsis '...' to expand the second slice
	combinedNumbers := append(numbers, moreNumbers...)
	fmt.Println("Combined slices:", combinedNumbers)
}

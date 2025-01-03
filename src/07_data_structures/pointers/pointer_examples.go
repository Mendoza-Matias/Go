package main

import "fmt"

// Define a struct to represent a Person
type Person struct {
	name  string
	age   int
	email string
}

// Method associated with the Person struct to print a greeting
func (p *Person) hello() {
	fmt.Println("Hello, my name is", p.name)
}

func main() {
	// Declaring an integer variable and its pointer
	var x int = 10
	var p *int = &x // Pointer that holds the memory address of x

	// Print the value of x
	fmt.Println("Value of x:", x)

	// Pass the reference of x to the edit function to modify it
	edit(&x)

	// Print the pointer value (memory address)
	fmt.Println("Pointer p:", p)

	// Print the updated value of x after modification
	fmt.Println("Updated value of x:", x)

	// Create an instance of Person and initialize it
	person := Person{"Matias", 22, "email@gmail.com"}
	// Call the hello method of the Person instance
	person.hello()
}

// Function to modify the value of x through a pointer
func edit(x *int) {
	*x = 20 // Dereference the pointer and modify the value of x
}

package main

import "fmt"

// Define a Person struct with fields: name, age, and email
type Person struct { 
	name  string
	age   int
	email string
}

func main() {

	// Create an instance of Person using the default struct constructor
	var p Person 
	p.name = "Matias"
	p.age = 22
	p.email = "email@gmail.com"

	// Create another instance of Person, initializing in one line
	p2 := Person{"Juan", 25, "juan@gmail.com"}

	// Print the entire Person structs
	fmt.Println(p)   // Output: {Matias 22 email@gmail.com}
	fmt.Println(p2)  // Output: {Juan 25 juan@gmail.com}

	// Print a specific field of the struct
	fmt.Println(p.name) // Output: Matias
}

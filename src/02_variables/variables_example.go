// First Variables Example
package main

import (
	"fmt"
)

func main() {

	// 1° Declaring multiple variables with `var` and assigning later
	var firstName, lastName string
	var age int

	// 2° Grouped variable declaration with initialization
	var (
		nationality string = "Argentine"
		address     string = "Montes de Oca"
		number      int    = 4746
	)

	// 3° Short variable declaration inside a function (type inferred)
	favoriteGame, favoriteFilm := "Call of Duty", "Coraline"

	// Reassigning a value to one of the variables
	favoriteGame = "Hitman"

	// Assigning values to variables declared earlier
	firstName = "Matias"
	lastName = "Mendoza"
	age = 21

	// Printing values
	fmt.Println(firstName, lastName, age)
	fmt.Println(nationality, address, number)
	fmt.Println(favoriteGame, favoriteFilm)
}

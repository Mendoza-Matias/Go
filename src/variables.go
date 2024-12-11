// first variables
package main

import(
	"fmt"
)

func main(){

	// 1°
	var firstName , lastName string 
	var age int

	// 2°
	var(
		nationality string = "Argentino"
		address string = "Montes de oca"
		number int = 4746
	)
  
	// 3° 
	
	//within the functions
	favoriteGame , favoriteFilm := "Call of Duty" , "Coroline"
	favoriteGame = "Hitman"
	

	firstName = "Matias"
	lastName = "Mendoza"
	age = 21

	fmt.Println(firstName, lastName, age)
	fmt.Println(nationality, address, number)
	fmt.Println(favoriteGame, favoriteFilm)
}
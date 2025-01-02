package main

import "fmt"

type Person struct { //model of a class
	name  string
	age   int
	email string
} //create structure

func main() {

	var p Person // create instance

	p.name = "Matias"
	p.age = 22
	p.email = "email@gmail.com"

	//create instance in one line

	p2 := Person{"Juan", 25, "juan@gmail.com"}

	fmt.Println(p)
	fmt.Println(p2)

	fmt.Println(p.name) // return Matias
}

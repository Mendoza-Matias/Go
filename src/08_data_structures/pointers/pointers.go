package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

/* method */
func (p *Person) hello() {
	fmt.Println("Hello my name is ", p.name)
}

func main() {
	var x int = 10
	var p *int = &x //pointer

	fmt.Println(x)

	edit(&x) //send reference of memory

	fmt.Println(p)

	fmt.Println(x)

	//instance Person

	person := Person{"Matias", 22, "email@gmail.com"}
	person.hello()
}

func edit(x *int) {
	*x = 20 //modify pointer
}

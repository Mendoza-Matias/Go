// first const
package main

import (
	"fmt"
)

// package level
const Pi float32 = 3.14 //declare and initialize
const (
	x int = 100
)
const (
	/*
		iota is used to initialize
		a constant or variable to
		0 and perform sequences
	*/
	Lunes     int = iota + 1 // 0 + 1 = 1
	Martes                   // 2
	Miercoles                // 3
	Jueves                   // 4
	Viernes                  // 5
)

func main() {
	const name string = "Hector"

	fmt.Println(Pi)
	fmt.Println(Martes)
	fmt.Println(name)
}

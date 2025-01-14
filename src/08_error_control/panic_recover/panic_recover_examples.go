package main

import (
	"fmt"
)

func main() {

	/* generate a panic */
	split(100, 10)
	split(200, 25)
	split(34, 0)
}
func split(dividend, divider int) {
	defer func() { /*

			USE IN EXCEPTIONAL CASES

			control panic and continue program flow
			at the end it throws the panic value */
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divider)
	fmt.Println(dividend / divider)
}

func validateZero(divider int) {
	if divider == 0 {
		panic("cannot be divided by zero")
	}
}

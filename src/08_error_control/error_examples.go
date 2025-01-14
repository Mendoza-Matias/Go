package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	str := "123"

	number, error := strconv.Atoi(str)

	if error != nil {
		fmt.Println("Error: ", error)
		return
	}
	fmt.Println(number)

	result, err := split(10, 2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	/* Use defer */

	defer fmt.Println(2) /* It runs at the end */
	fmt.Println(1)
	fmt.Println(3)

	file, e := os.Create("file.txt")

	if e != nil {
		fmt.Println(e)
		return
	}

	defer file.Close() /* It is executed before the main function finishes executing. */

	_, e = file.Read([]byte("Hello"))

	if e != nil {
		fmt.Println(e)
		return
	}
}

func split(dividend, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("cannot be divided by zero") /* create a error */
	}
	return dividend / divider, nil
}

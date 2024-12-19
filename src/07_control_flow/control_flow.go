package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//conditional

	// if condicion {

	// }else {

	// }

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("tomorrow")
	} else if t.Hour() < 17 {
		fmt.Println("late")
	} else {
		fmt.Println("night")
	}
	/* initialize the variable to use only in the if */

	// switch value {
	// 	case
	// }

	os := runtime.GOOS /* operating system */

	switch os {
	case "windows":
		fmt.Println("go run windows")
	case "linux":
		fmt.Println("go run linux")
	default:
		fmt.Println("go run other os")
	}

	switch num := 10; num {
	case 10:
		fmt.Println("ten")
	case 7:
		fmt.Println("seven")
	default:
		fmt.Println("other number")
	}

	//------FIRST LOOP---------------

	// for condition {
	// }

	var i int

	for i <= 10 {
		fmt.Print(" ", i)
		i++
		if i == 6 {
			break /*stop loop*/
		}
	}

	fmt.Println("")

	for i := 0; i <= 6; i++ {
		if equal2(i) {
			continue
		}
		fmt.Print(" ", i) // 0 1 3 4 5 6
	}

	fmt.Println("")

	s, r := twoResult(5, 2) /* capture two result in variables */

	fmt.Println("result =>", s)
	fmt.Println("result =>", r)

}

/* first function */
func equal2(number int) bool {
	return number == 2
}

/*
return two values

not recommended
better : return a value , a value and error or nothing
*/
func twoResult(n1 int, n2 int) (s, r int) {
	s = n1 + n2
	r = n1 - n2
	return s, r
}

package main //name of package

import (
	"fmt"
	 "rsc.io/quote"
	) // external package

func main() {
	fmt.Println("Hello world")
	fmt.Println(quote.Hello())
}

package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func main() {

	var sideA, sideB float64

	fmt.Print("enter the first side ")
	fmt.Scan(&sideA)
	fmt.Print("enter the second side ")
	fmt.Scan(&sideB)

	sum := (math.Pow(sideA, 2) + math.Pow(sideB, 2))

	result := math.Sqrt(sum)

	fmt.Println(decimal.NewFromFloat(result).Round(2))
	fmt.Println("area", (sideA+sideB)/2)

}

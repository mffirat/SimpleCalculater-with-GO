package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64
	var num2 float64
	var operate int
	var result float64

	fmt.Println("Enter num1: ")
	fmt.Scan(&num1)

	fmt.Println("Enter num2: ")
	fmt.Scan(&num2)

	fmt.Print("Choose a operater:\n Enter 1 for addition:\n Enter 2 for subtraction:\n Enter 3 for multiplication:\n Enter 4 for division:\n Enter 5 for Exponentiation\n Enter 6 for root\n")
	fmt.Scan(&operate)

	if operate == 1 {
		result = num1 + num2
	} else if operate == 2 {
		result = num1 - num2
	} else if operate == 3 {
		result = num1 * num2
	} else if operate == 4 {
		if num2 == 0 {
			fmt.Print("A number does not divide with 0")
		} else {
			result = num1 / num2
		}
	} else if operate == 5 {
		result = math.Pow(num1, num2)
	} else if operate == 6 {
		if num2 == 0 {
			fmt.Print("A number does not root with 0")
		} else {
			result = math.Pow(num1, 1/num2)
		}
	} else {
		fmt.Print("PLS choose a valid operate\n")
	}
	fmt.Printf("Result: %f", result)

}

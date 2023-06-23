package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func main() {
	var x int
	fmt.Println("Type a number ")
	fmt.Scanln(&x)
	x = int(x)

	var operation string
	fmt.Println("Choose an operator")
	fmt.Scanln(&operation)
	operation = string(operation)

	var y int
	fmt.Println("Type an other number ")
	fmt.Scanln(&y)
	y = int(y)

	if operation == "+" {
		result := (sum(x, y))
		fmt.Println("Your result is", result)
	} else {
		if operation == "-" {
			result := (sub(x, y))
			fmt.Println("Your result is", result)
		} else {
			if operation == "*" {
				result := (mul(x, y))
				fmt.Println("Your result is", result)
			} else {
				if operation == "/" && y != 0 {
					result := (div(x, y))
					fmt.Println("Your result is", result)
				} else {
					fmt.Println("Operation not supported")
				}
			}

		}
	}

}

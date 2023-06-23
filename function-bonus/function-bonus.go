package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	var x int
	fmt.Println("Type a number ")
	fmt.Scanln(&x)
	x = int(x)

	var y int
	fmt.Println("Type an other number ")
	fmt.Scanln(&y)
	y = int(y)

	result := (sum(x, y))
	fmt.Println("Your result is", result)
}

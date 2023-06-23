package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	var x int
	fmt.Println("Type a numer ")
	fmt.Scanln(&x)
	x = int(x)

	var y int
	fmt.Println("Type an other numer ")
	fmt.Scanln(&y)
	y = int(y)

	result := (sum(x, y))
	fmt.Println("Your result is", result)
}

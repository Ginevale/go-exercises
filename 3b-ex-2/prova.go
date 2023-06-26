package main

import "fmt"

func main() {
	a := []int{1, 2}
	a = append(a, 3, 4) // a == [1 2 3 4]
	fmt.Println(a)
}

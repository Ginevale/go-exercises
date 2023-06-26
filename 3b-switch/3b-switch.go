package main

import (
	"errors"
	"fmt"
	"log"
)

func sum(a, b int) int {
	return a + b
}

func diff(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Operation not supported")
	} else {
		return a / b, nil
	}
}

func res(array []int, operator string, result int) {
	for i := 0; i < len(array)-1; i++ {
		fmt.Print(array[i], " ", operator, " ")
	}
	fmt.Println(array[len(array)-1], "=", result)
}

func main() {
	var rep int
	fmt.Println("How many numbers do you want to use?")
	fmt.Scanln(&rep)
	rep = int(rep)
	if rep == 0 {
		fmt.Println("No operation required")
		return
	}

	var operation string
	fmt.Println("Choose an operator")
	fmt.Scanln(&operation)
	operation = string(operation)

	var nums []int

	for i := 0; i < rep; i++ {
		var in int
		fmt.Println("Type a number ")
		fmt.Scanln(&in)
		in = int(in)
		nums = append(nums, in)
	}

	total := nums[0]

	for i := 1; i < rep; i++ {
		switch {
		case operation == "+":
			total = sum(total, nums[i])

		case operation == "-":
			total = diff(total, nums[i])

		case operation == "*":
			total = mul(total, nums[i])

		case operation == "/":
			res, err := div(total, nums[i])
			if err != nil {
				log.Fatal(err)
			}
			total = res
		}
		res(nums, operation, total)
	}
}

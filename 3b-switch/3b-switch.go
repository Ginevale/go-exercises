package main

import "fmt"

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

	switch {
	case operation == "+":
		for i := 1; i < rep; i++ {
			total += nums[i]
		}

	case operation == "-":
		for i := 1; i < rep; i++ {
			total -= nums[i]
		}

	case operation == "*":
		for i := 1; i < rep; i++ {
			total *= nums[i]
		}

	case operation == "/":
		for i := 1; i < rep; i++ {
			if nums[i] == 0 {
				fmt.Println("Operation not supported")
				return
			} else {
				total /= nums[i]
			}
		}

	default:
		fmt.Println("Operation not supported")
		return
	}
	res(nums, operation, total)
}

package main

import "fmt"

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

	var first int
	fmt.Println("Type the first number")
	fmt.Scanln(&first)
	first = int(first)
	nums = append(nums, first)

	for i := 1; i <= rep-1; i++ {
		var in int
		fmt.Println("Type a number ")
		fmt.Scanln(&in)
		in = int(in)
		nums = append(nums, in)
	}

	if operation == "+" {
		var sum int
		for i := 0; i <= rep-1; i++ {
			sum += nums[i]
		}
		for i := 0; i <= rep-2; i++ {
			fmt.Print(nums[i], " ", operation, " ")
		}
		fmt.Println(nums[rep-1], "=", sum)

	} else {
		if operation == "-" {
			diff := first
			for i := 1; i <= rep-1; i++ {
				diff -= nums[i]
			}
			for i := 0; i <= rep-2; i++ {
				fmt.Print(nums[i], " ", operation, " ")
			}
			fmt.Println(nums[rep-1], "=", diff)

		} else {
			if operation == "*" {
				mul := first
				for i := 1; i <= rep-1; i++ {
					mul *= nums[i]
				}
				for i := 0; i <= rep-2; i++ {
					fmt.Print(nums[i], " ", operation, " ")
				}
				fmt.Println(nums[rep-1], "=", mul)

			} else {
				if operation == "/" {
					div := first
					for i := 1; i < rep; i++ {
						if nums[i] == 0 {
							fmt.Println("Operation not supported")
							return
						} else {
							div /= nums[i]
						}
					}
					for i := 0; i <= rep-2; i++ {
						fmt.Print(nums[i], " ", operation, " ")
					}
					fmt.Println(nums[rep-1], "=", div)

				} else {
					fmt.Println("Operation not supported")
					return
				}
			}

		}
	}

}

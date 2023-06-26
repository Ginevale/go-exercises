package main

import "fmt"

func main() {
	var rep int
	fmt.Println("How many numbers do you want to use?")
	fmt.Scanln(&rep)
	rep = int(rep)

	var operation string
	fmt.Println("Choose an operator")
	fmt.Scanln(&operation)
	operation = string(operation)

	var first int
	fmt.Println("Type the first number ")
	fmt.Scanln(&first)
	first = int(first)

	if operation == "+" {
		sum := first
		for i := 2; i <= rep; i++ {
			var in int
			fmt.Println("Type a number ")
			fmt.Scanln(&in)
			in = int(in)
			sum += in
		}
		fmt.Println("Your result is", sum)

	} else {
		if operation == "-" {
			diff := first
			for i := 2; i <= rep; i++ {
				var in int
				fmt.Println("Type a number ")
				fmt.Scanln(&in)
				in = int(in)
				diff -= in
			}
			fmt.Println("Your result is", diff)

		} else {
			if operation == "*" {
				mul := first
				for i := 2; i <= rep; i++ {
					var in int
					fmt.Println("Type a number ")
					fmt.Scanln(&in)
					in = int(in)
					mul *= in
				}
				fmt.Println("Your result is", mul)

			} else {
				if operation == "/" {
					div := first
					for i := 2; i <= rep; i++ {
						var in int
						fmt.Println("Type a number ")
						fmt.Scanln(&in)
						in = int(in)
						if in == 0 {
							fmt.Println("Operation not supported")
						} else {
							div /= in
						}
					}
					fmt.Println("Your result is", div)

				} else {
					fmt.Println("Operation not supported")
				}
			}

		}
	}

}

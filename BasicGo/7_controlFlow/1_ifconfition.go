package main

import "fmt"

func Ifelseifelse() {
	number := 0

	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	} else if number < 0 {
		fmt.Printf("%d is a negative number\n", number)
	} else {
		fmt.Printf("%d is neither positive nor negative\n", number)
	}
}

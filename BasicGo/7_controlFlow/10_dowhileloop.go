package main

import "fmt"

func DoWhileLoop() {
	number := 3
	// loop that runs infinitely
	for {
		// condition to terminate the loop
		if number > 5 {
			break
		}
		fmt.Printf("%d\n", number)
		number++
	}
}

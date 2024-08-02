package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// call the function
	go printNumbers()

	for i := 'a'; i <= 'e'; i++ {
		fmt.Println(string(i))
		time.Sleep(1 * time.Second)
	}

	// sleep for 6 seconds
	time.Sleep(6 * time.Second)
}

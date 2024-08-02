package main

import "fmt"

func SwitchMultiple() {
	dayOfWeek := "Sunday"
	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid day")
	}
}

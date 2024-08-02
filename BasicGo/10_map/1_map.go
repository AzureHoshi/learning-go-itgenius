package main

import "fmt"

func main() {
	// Crate Map
	subjectMarks := map[string]float32{
		"Golang": 85.5,
		"Python": 90.0,
		"Java":   80.0,
	}
	fmt.Println(subjectMarks)
	fmt.Println(subjectMarks["Golang"])
	fmt.Println(subjectMarks["Python"])
	fmt.Println(subjectMarks["Java"])
	// Change value
	subjectMarks["Java"] = 88.0
	fmt.Println(subjectMarks["Java"])
	// Add new key
	subjectMarks["C++"] = 75.0
	fmt.Println(subjectMarks)
	// Delete key
	delete(subjectMarks, "C++")
	fmt.Println(subjectMarks)
}

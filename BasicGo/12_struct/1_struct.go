package main

import "fmt"

func main() {

	// Declare a struct
	type Person struct {
		name string
		age  int
	}

	// assign value to struct
	person1 := Person{"John", 30}
	fmt.Println(person1) // {John 30}

	//  define a struct with field names
	var person2 Person

	//  assign value to struct with field names
	person2 = Person{name: "Doe", age: 25}

	fmt.Println(person2) // {Doe 25}

}

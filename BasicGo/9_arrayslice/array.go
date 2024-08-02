package main

import "fmt"

func CreateArray() {

	// declare array variable of type integer
	// defined size [size=5]
	var arrayOfInteger = [5]int{1, 5, 8, 0, 3}

	// undefined size
	var arrayOfInteger2 = [...]int{1, 5, 8, 0, 3, 4, 6, 7, 8, 9}

	// using shorthand notation
	arrayOfInteger3 := [5]int{1, 5, 8, 0, 3}

	fmt.Println(arrayOfInteger)
	fmt.Println(arrayOfInteger2)
	fmt.Println(arrayOfInteger3)
}

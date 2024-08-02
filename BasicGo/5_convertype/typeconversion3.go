// แปลงจากจำนวนเต็มเป็นสตริง

package main

import (
	"fmt"
	"strconv"
)

func typeConversion3() {
	fmt.Println("=== Type Conversion 3 ===")
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("i: %d, s: %s\n", i, s)
	// ผลลัพธ์: i: 42, s: 42
}

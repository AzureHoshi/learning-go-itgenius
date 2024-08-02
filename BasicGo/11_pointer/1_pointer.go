package main

func main() {
	i, j := 42, 2701

	p := &i     // point to i
	println(*p) // read i through the pointer

	*p = 21 // set i through the pointer

	println(i) // see the new value of i

	p = &j // point to j
	println(*p)

	*p = *p / 37 // divide j through the pointer

	println(j) // see the new value of j

	// 42
	// 21
	// 73
}

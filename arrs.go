package main

import "fmt"

func main() {
	// array with size 2 type string. cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	// prints without , between elements.
	fmt.Println(a)

	// array with size 6 type int
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}


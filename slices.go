package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// its called a slice. similar to numpy (?)
	// don't know the use yet..
	var s []int = primes[1:4]
	fmt.Println(s)
}


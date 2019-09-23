package main

import "fmt"

func main() {
	// doesn't run until function ends.
	defer fmt.Println("world")

	fmt.Println("hello")
}

package main

import "fmt"

func main() {
	var s []int
	// nil
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

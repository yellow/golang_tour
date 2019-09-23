package main

import "fmt"

func main() {
	sum := 0
	// for doesn't require parentheses.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}


package main

import "fmt"

// functions start with keyword func. datatypes after variable names.
// function return types at the end. again, no semi-colons ;-; 
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

  // a points to the first 2 elements in array
	a := names[0:2]
  // b points to the last 2 elements in array
	b := names[1:3]
	fmt.Println(a, b)

  // change in b means ...
	b[0] = "XXX"
	fmt.Println(a, b)
  // change in original array.
	fmt.Println(names)
}


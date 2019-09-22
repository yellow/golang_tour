package main

import "fmt"

// returning two strings below
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
  // no need to declare "hello" and "world" as string because of ':='
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

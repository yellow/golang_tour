package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

  // starting 0 is default
	s = s[:2]
	fmt.Println(s)

  // till the end
	s = s[1:]
	fmt.Println(s)
}

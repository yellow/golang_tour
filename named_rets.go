package main

import "fmt"

// they have named the return ints to x and y
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
  // naked return. can harm readability
	return
}

func main() {
	fmt.Println(split(17))
}


package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
  // set value to v and compare. v only in scope of if statement.
	if v := math.Pow(x, n); v < lim {
		return v
	}
  // no v here.
	return lim
}

func main() {
  // one way to print...
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

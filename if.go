package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// if doesn't need () too.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// pretty print maybe..?
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}


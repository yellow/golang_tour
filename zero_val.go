package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
  // outputs blank as its garbage collected.
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

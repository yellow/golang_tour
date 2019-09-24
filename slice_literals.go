package main

import "fmt"

func main() {
  // slice literal with values and no number of elements.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

  // array of structs
	s := []struct {
		i int
		b bool
	}{
    // multiple structs defined inside the array
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}


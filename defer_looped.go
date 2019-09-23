package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// defer looped. gets added to stack. 9 8 7 6 5
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

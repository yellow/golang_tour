package main

import "fmt"

func main() {
	sum := 1
  // for used like a while. without the semicolon. looks better
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}


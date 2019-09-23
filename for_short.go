package main

import "fmt"

func main() {
	sum := 1
	// for without initialization and operation.
	// btw golang has no while loop / do while / etc.
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}


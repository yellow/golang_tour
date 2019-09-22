package main

import "fmt"

// declared and initialized
var i, j int = 1, 2

func main() {
  // same
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}


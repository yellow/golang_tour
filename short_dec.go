package main

// another way to import. not recommended
import "fmt"

func main() {
  // initialized
	var i, j int = 1, 2
  // declared without datatype
	k := 3
  // same with multiple vars and different types
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

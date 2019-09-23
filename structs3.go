package main

import "fmt"

// struct Vertex declared.
type Vertex struct {
	X int
	Y int
}

func main() {
	// v initialized
	v := Vertex{1, 2}
	// gets address of struct v and fills p with it.
	p := &v
  // this still works!. ( (*p).X will too. )
	p.X = 1e9
	fmt.Println(v)
}

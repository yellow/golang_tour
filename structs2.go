package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// dot operator to get data from struct v
	v.X = 4
	fmt.Println(v.X)
}

package main

import (
  "fmt"
  "math/rand"
)

func main() {
  // same number because no seed.
  fmt.Println("Random number: ", rand.Intn(10))
}

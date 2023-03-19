package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4  // type inference used for x and y
	lucky := math.Sqrt(float64(x*x + y*y))  // defining variabe  
	z := uint(lucky)  // type conversion simplified
	fmt.Println(x, y, z)
}

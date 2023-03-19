package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	toBe := true
	maxInt := uint64(1<<63 - 1) // decreased by 1 to avoid overflow
	z := cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

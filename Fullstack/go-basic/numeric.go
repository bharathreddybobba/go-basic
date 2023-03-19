package main

import "fmt"

func main() {
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Printf("Small is of type %T and has a value of %d\n", Small, Small)
	fmt.Printf("Big is of type %T and has a value of %d\n", Big, Big)

	var v1 int = int(Small)
	var v2 float64 = float64(Small)
	var v3 float64 = float64(Big)

	fmt.Println(needInt(v1))
	fmt.Println(needFloat(v2))
	fmt.Println(needFloat(v3))
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

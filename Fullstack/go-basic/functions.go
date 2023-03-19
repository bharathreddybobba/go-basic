package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	
	fmt.Print("Enter two integers: ")

	
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Printf("The sum of %d and %d is %d\n", x, y, add(x, y))
}

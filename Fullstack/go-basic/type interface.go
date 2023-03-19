package main
import "fmt"

func main() {
	var x int = 10
	var y float64 = 3.14
	var z string = "Hello, world!"

	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)

	var input interface{}
	fmt.Println("Enter a value:")
	fmt.Scan(&input)
	fmt.Printf("You entered a value of type %T\n", input)
}

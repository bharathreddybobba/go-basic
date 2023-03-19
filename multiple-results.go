package main

import "fmt"

func swap(s string) string {
	return s[1:] + string(s[0])
}

func main() {
	s := "hello"
	s = swap(s)
	fmt.Println(s)
}

package main

import (
	"fmt"
	"unicode"
)

func isCamelCase(s string) bool {
	if len(s) == 0 {
		return false
	}

	if unicode.IsUpper(rune(s[0])) {
		return false
	}

	for _, r := range s[1:] {
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			return false
		}

	}

	return true
}

func main() {
	str1 := "helloWorld"
	str2 := "HelloWorld"
	str3 := "hello_world"
	str4 := ""

	fmt.Printf("%s is camel case: %t\n", str1, isCamelCase(str1))
	fmt.Printf("%s is camel case: %t\n", str2, isCamelCase(str2))
	fmt.Printf("%s is camel case: %t\n", str3, isCamelCase(str3))
	fmt.Printf("%s is camel case: %t\n", str4, isCamelCase(str4))
}

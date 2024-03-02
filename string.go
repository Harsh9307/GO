package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, world!"
	var subStr string = "  wor"

	if strings.Contains(str, subStr) == true {
		fmt.Printf("String (%s) contains sub-string (%s)", str, subStr)
	} else {
		fmt.Printf("String (%s) does not contain sub-string (%s)\n", str, subStr)
	}
	var result string
	result = strings.ToUpper(str)
	fmt.Println("String in uppercase:", result)

	var ind int = 0
	ind = strings.Index(str, "w")
	fmt.Println("Index is :", ind)

	// Replace part of a string
	newStr := strings.Replace(str, "world", "Go", -1)
	fmt.Println("String after replacement:", newStr)

	// Trim leading and trailing whitespaces
	trimmedStr := strings.TrimSpace("   Hello, world!   ")
	fmt.Println("String after trimming spaces:", trimmedStr)

	// Split a string into substrings based on a delimiter
	splitStr := strings.Split("apple,orange,banana", ",")
	fmt.Println("String split into substrings:", splitStr)

	// Join substrings into a single string
	joinedStr := strings.Join(splitStr, "-")
	fmt.Println("Substrings joined into a single string:", joinedStr)
}

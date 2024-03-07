package main

import "fmt"

func main() {
	a := 10
	p := &a

	fmt.Println("Printing Address...")
	fmt.Println(a, &a)
	fmt.Println(a, p)

	fmt.Println("Printing values...")
	fmt.Println(a, *(&a))
	fmt.Println(a, *p)
}

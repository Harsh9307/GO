/*
Go Slice is an abstraction over Go Array
Go Array allows you to defune variables that can go
*/
package main

import "fmt"

func main() {
	//this is array
	//number := [5]int{1, 2, 3, 4, 5}

	//this is slice
	numbers := []int{1, 2, 3, 4, 5}
	num1 := numbers[:4]
	fmt.Println(num1)

	num2 := numbers[0:]
	fmt.Println(num2)

	fmt.Println("Numbers", numbers)

}

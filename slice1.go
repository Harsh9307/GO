package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
}
func printSlice(numbers []int) {
	fmt.Printf("len = %d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}

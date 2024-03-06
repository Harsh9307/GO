package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

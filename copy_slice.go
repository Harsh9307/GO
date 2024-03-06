package main

import "fmt"

func main() {
	//create 2 slices
	primeNumbrs := []int{2, 3, 5, 7}
	numbers := []int{1, 2, 3}

	copy(numbers, primeNumbrs)
	fmt.Println("Numbers:", numbers)
}

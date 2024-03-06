package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{70, 20, 30, 50, 60, 10, 80, 90, 100}
	sort.Ints(slice)

	fmt.Println("Sorted Slice:")
	for _, ele := range slice {
		fmt.Println(ele)
	}
}

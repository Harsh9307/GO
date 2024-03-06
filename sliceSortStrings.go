package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"honesty", "is", "the", "best", "policy"}

	sort.Strings(slice)
	for _, item := range slice {
		fmt.Printf("%s ", item)
	}
}

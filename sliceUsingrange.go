package main

import "fmt"

func main() {
	// create an integer array
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//create slice of from index 2 till index 4
	intSlice := arr[2:5]

	fmt.Println("Slice Elements:")
	for index, ele := range intSlice {
		fmt.Printf("index: %d, element: %d\n", index, ele)
	}
	for _, ele := range intSlice {
		fmt.Println(ele)
	}
}

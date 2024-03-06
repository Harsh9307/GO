package main

import (
	"fmt"
)

func isSorted(slice1 []int) {

	for i := 0; i < len(slice1)-1; i++ {
		if slice1[i] > slice1[i+1] {
			fmt.Println("Slice 1 is not sorted")
			return
		}
	}
	fmt.Println("Slice is sorted")

}
func main() {
	//var status bool = false
	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	//slice2 := []int{70, 20, 30, 60, 50, 60, 10, 80, 90, 100}

	isSorted(slice1)
	// if status == true {
	// 	fmt.Println("Slice 1 is sorted")
	// } else {
	// 	fmt.Println("Slice is not sorted")
	// }
	// status = sort.IntsAreSorted(slice2)
	// if status == true {
	// 	fmt.Println("Slice 1 is sorted")
	// } else {
	// 	fmt.Println("Slice 2 is not sorted")
	// }
}

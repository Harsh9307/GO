package main

import "fmt"

func main() {
	// x := 8
	// y := 10
	// if x != y {
	// 	if x < y {
	// 		fmt.Println("x is less than y")
	// 	} else if x > y {
	// 		fmt.Println("x is greater than y")
	// 	} else {
	// 		fmt.Println("x is equal to y")
	// 	}
	// }
	var x int
	fmt.Println("Enter the time")
	fmt.Scanf("%d", &x)

	if x <= 10 {
		fmt.Println("Good Morning ☀️")
	} else if x > 10 && x <= 20 {
		fmt.Println("Good Day :)")
	} else {
		fmt.Println("Good Evening :(")
	}
}

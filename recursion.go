package main

import "fmt"

func sum(number int) int {
	// condition to break recursion
	if number <= 0 {
		return 0
	} else {
		return number + sum(number-1)
	}
}
func main() {
	var num int
	fmt.Println("Enter the number :")
	fmt.Scanf("%d", &num)
	//function call
	var result = sum(num)
	fmt.Println("Sum: ", result)

}

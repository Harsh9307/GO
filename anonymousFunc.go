package main

import "fmt"

// regular function to calculate the square of numbers
func findSquare(num int) int {
	square := num * num
	return square
}

func main() {
	// anonymous function
	// var greet = func() {
	// 	fmt.Println("Hello , how are you")
	// }
	// welcome := greet
	// //function call
	// greet()
	// welcome()

	// //with parameters
	// var add = func(n1 int, n2 int) {
	// 	sum := n1 + n2
	// 	fmt.Println("Sum is : ", sum)
	// }
	// add(5, 3)
	// fmt.Printf("the typ of add is : %T", add)

	// var sum = func(n1 int ,n2 int) int{
	// 	sum :=n1+n2
	// 	return sum
	// }

	// //function call
	// result := sum(5,3)
	// fmt.Println("sum is : ", result)

	// annoymous function that returns sum of  numbers
	sum := func(number1 int, number2 int) int {
		return number1 + number2
	}

	//function call
	result := findSquare(sum(6, 9))
	fmt.Println("result is : ", result)

}

// Fibonacci series
package main

import "fmt"

func main() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))


	// **************cannot do like this***********************
	// //var fib func(n int) int
	// fib1 := func(n int) int {
	// 	if n < 2 {
	// 		return n
	// 	}
	// 	return fib1(n-1) + fib1(n-2)
	// }
	// fmt.Println(fib1(7))
}

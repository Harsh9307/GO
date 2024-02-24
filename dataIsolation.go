package main

import "fmt"

func displayNumbers() func() int {
	number := 0
	//inner function
	return func() int {
		number++
		return number
	}
}
func main() {
	//returns a closure
	num1 := displayNumbers()
	fmt.Printf("The type of num1 is %T\n", num1)
	fmt.Println(displayNumbers()()) //1
	fmt.Println(displayNumbers()()) //1
	fmt.Println()
	fmt.Println(num1()) //1
	fmt.Println(num1()) //2
	fmt.Println(num1()) //3

	//returns a new closure
	num2 := displayNumbers()
	fmt.Println(num2())
	fmt.Println(num2())

	//num1 does not get destroyed even after num2 is called
	fmt.Println(num1())

	//it will return the same value... reinitialization is done but if function returns a closure .then the increment will be done
	add := func() int {
		num := 1
		return num + 1
	}
	add1 := add
	fmt.Println(add1())
	fmt.Println(add1())
	fmt.Println(add1())

}

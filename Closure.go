package main

import "fmt"

// func calculate() func() int {
// 	num := 1
// 	//return inner function
// 	return func() int {
// 		num = num + 2
// 		return num
// 	}

// }
func greet() func() string {
	// variable defined outside the inner function
	name := "Harsh"
	// return a nested annoymous function
	return func() string {
		name := "Hii" + " " + name + "\n"
		return name
	}
}
func incrementNumber() func() int {
	num := 1
	return func() int {
		num += 1
		return num
	}
}
func main() {
	// //call the outer function
	// odd := calculate()
	// // call the inner function
	// fmt.Println(odd())   //3
	// fmt.Println(odd())	//5
	// fmt.Println(odd())   //7

	// // call the outer function again
	// odd2 := calculate()
	// fmt.Println(odd2())        //3

	//call the outer function again
	message := greet()
	//call the inner function again
	fmt.Println(message())
	fmt.Print(message())

	inc := incrementNumber()
	fmt.Println(inc())
	fmt.Println(inc())

}

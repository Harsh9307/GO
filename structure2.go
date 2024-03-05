package main

import "fmt"

// Define a struct named Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Function that takes a Person struct as a parameter
func printPersonInfo(p Person) {
	fmt.Printf("Name: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
}

func main() {
	// Create an instance of the Person struct
	john := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Call the function with the Person instance
	printPersonInfo(john)
}

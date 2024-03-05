package main

import "fmt"

// Define a function type for calculating the area of a rectangle
type Rectangle func(int, int) int

// Define a struct to represent a rectangle with additional properties
type RectanglePara struct {
	Length  int
	Breadth int
	Color   string

	// Use the Rectangle function type for the rect field
	rect Rectangle
}

func main() {
	// Create an instance of RectanglePara
	result := RectanglePara{
		Length:  10,
		Breadth: 10,
		Color:   "Red",
		// Initialize the Rect field with a function literal
		rect: func(length, breadth int) int {
			return length * breadth
		},
	}

	// Print information about the rectangle
	fmt.Println("Color of Rectangle:", result.Color)
	fmt.Println("Area of Rectangle:", result.rect(result.Length, result.Breadth))
}

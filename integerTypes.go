package main

import "fmt"

func main() {
	var a uint8
	a = 10
	var student1 string = "John" // type is string
	var student2 = "Jane"        // type is inferred
	x := 2                       // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(a)

	fmt.Print(a, student1, student2)
	fmt.Printf("\n")
	fmt.Println(a, student1, student2)

	var x1 float64 = 20.0
	y := 42
	fmt.Printf("x = %f\n", x1)
	fmt.Printf("y = %d\n", y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

	var a1 string
	var b int
	var c bool

	c = true
	fmt.Println(a1)
	fmt.Println(b)
	fmt.Println(c)

	var d, e, f = 3, 4, "foo"
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}

package main

import "fmt"

const (
	A int = 1
	B     = 3.14
	C     = "Hii"
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 20
	const Const int = 20


	fmt.Println(A, B, C)
	var area int
	area = LENGTH * WIDTH
	fmt.Printf("The area is %d", area)
}

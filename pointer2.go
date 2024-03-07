package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	t := *x
	*x = *y
	*y = t
}
func main() {

	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}

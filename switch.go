package main

import (
	"fmt"
)

func main() {
	// thisMonth := 5
	// switch thisMonth {
	// case 1:
	// 	fmt.Println("January")
	// case 2:
	// 	fmt.Println("February")
	// case 3:
	// 	fmt.Println("March")
	// case 4:
	// 	fmt.Println("April")
	// case 5:
	// 	fmt.Println("May")
	// }

	// today := time.Now()
	// switch today.Day() {
	// case 2:
	// 	fmt.Println("Today is Saturday")
	// 	fmt.Println("Today is", today.Format("2006-01-02 (Monday)"))
	// 	dayOfWeek := today.Weekday()
	// 	if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
	// 		fmt.Println("It's the weekend!")
	// 	} else {
	// 		fmt.Println("It's a weekday.")
	// 	}
	// 	currentHour := today.Hour()
	// 	if currentHour >= 9 && currentHour < 17 {
	// 		fmt.Println("It's working hours.")
	// 	} else {
	// 		fmt.Println("It's outside working hours.")
	// 	}

	// case 5:
	// 	fmt.Println("Today is Sunday")
	// default:
	// 	fmt.Println("No information available")
	// }

	// thisMonth := time.Now().Month()

	// switch thisMonth {
	// case 1, 3, 5, 7, 10, 12:
	// 	fmt.Println("31 days")
	// case 4, 6, 8, 11:
	// 	fmt.Println("30 days")
	// case 2:
	// 	fmt.Println("28 days")
	// }

	// x := 3
	// switch x {
	// case 1:
	// 	fmt.Println("1")
	// 	fallthrough
	// case 3:
	// 	fmt.Println("3")
	// 	fallthrough
	// case 5:
	// 	fmt.Println("5")
	// }

	//var x interface{}
	var x interface{} = true
	switch i := x.(type) {
	case nil:
		fmt.Printf("Type of x: %T", i)
	case int:
		fmt.Println("x is int")
	case float64:
		fmt.Println("x is float64")
	case func(int) float64:
		fmt.Printf("x is bool or string")
	case string:
		fmt.Println("x is string")
	default:
		fmt.Printf("dont know this type")
	}

}

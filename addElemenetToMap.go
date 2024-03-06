package main

import "fmt"

func main() {
	//create a map
	students := map[int]string{1: "John", 2: "Lily"}
	fmt.Println("Initial Map: ", students)

	//add element with key 3
	students[3] = "robin"

	students[4] = "Julice"
	students[7] = "a"
	students[6] = "b"

	fmt.Println("Updated Map: ", students)

	delete(students, 2)
	fmt.Println("Updated Map: ", students)

	for id, name := range students {
		fmt.Printf("id of %d is %s\n", id, name)
	}
}

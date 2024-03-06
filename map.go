package main

import "fmt"

func main() {
	//create a mao
	subjectMarks := map[string]float32{"GoLang": 85, "Java": 80, "Python": 81}
	fmt.Println(subjectMarks)
	for key, _ := range subjectMarks {
		fmt.Println(key)
	}
}

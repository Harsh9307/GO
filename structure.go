package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1, Book2 Books
	Book1.title = "Go Programming"
	Book1.author = "a"
	Book1.subject = "go"
	Book1.book_id = 12345

	Book2.title = "machine learning simplified"
	Book2.author = "b"
	Book2.subject = "machine learning"
	Book2.book_id = 98745

	fmt.Printf("Book1 title = %s\n", Book1.title)
	fmt.Printf("Book2 title = %s", Book2.title)

}

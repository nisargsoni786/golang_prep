package main

import "fmt"
import "reflect"

type Books struct{
	title string
	author string
	book_id int
}

func main() {
	var book1 Books
	var book2 Books

	book1.book_id=1
	book1.title="Story of my life"
	book1.author="one direction"

	// book2.book_id=2                   // this two way we can define struct variable
	// book2.title="You & I"
	book2=Books{book_id:2,title:"You & I"}

	fmt.Println(book1)
	fmt.Println(book1.author)
	fmt.Println()
	fmt.Println()
	fmt.Println(book2)
	fmt.Println(book2.author,reflect.TypeOf(book2.author))  // this will be none because we dont define author for the book2
	fmt.Println(book2)

}
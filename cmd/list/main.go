// Every Go source file begins with a 'package clause'.
// A package groups code together, allowing to use them as a single unit.
// You can name it whatever you want, but 'main' is special - compiling the main package produces an executable.
// Main package cannot be imported by other code. A test for the main package should use `package main` too
package main

// This file was moved into cmd/list:
// - "cmd" is just a convention, used for commands the user can run
// - this file lists books, so "lists" subdirectory
// go run ./cmd/list

// Declare we'll need to use this package too
import (
	"books"
	"fmt"
)

// // declare a slice variable
// var books []Book
// // initialize with a slice literal
// books := []Book{}

// Function declaration
func main() {
	// A statement.
	// Declaration = gives descriptive information about the program
	// Statement = indicates to perform an action
	//hello()

	catalog := books.GetCatalog()
	for _, book := range books.GetAllBooks(catalog) {
		fmt.Println(books.BookToString(book))
	}
}

func structsExample() {
	// single book
	book := books.Book{
		Title:  "Artemis Fowl",
		Author: "Eoin Colfer",
		Copies: 3_000_000,
	}

	fmt.Println(books.BookToString(book))

	poetryBook := books.Book{
		Title:  "Bolle di sapone",
		Author: "Gianni Rodari",
		Copies: 100_000,
	}

	fmt.Println(books.BookToString(poetryBook))
}

func slicesExample() {
	myBooks := []books.Book{
		{
			Title:  "Artemis Fowl",
			Author: "Eoin Colfer",
			Copies: 3_000_000,
		},
		{
			Title:  "Bolle di sapone",
			Author: "Gianni Rodari",
			Copies: 100_000,
		},
	}
	myBooks[1].Copies = 200_000
	fmt.Println(len(myBooks)) // 2
	// myBooks[2].Copies = 1 // runtime panic! out of bounds

	// ----- Adding an element -----
	newBook := books.Book{
		Title:  "Jujutsu Kaisen vol. 1",
		Author: "Gege Akutami",
		Copies: 10_000_000,
	}
	// `append` returns a new or modified slice with the added element
	myBooks = append(myBooks, newBook)
}

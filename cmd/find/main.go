package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	// First argument to os.Args is always the name of the program
	if len(os.Args) < 2 {
		fmt.Println("Usage: find <book_id>")
		return
	}

	catalog := books.GetCatalog()
	bookID := os.Args[1]
	book, ok := books.GetBook(catalog, bookID)
	if !ok {
		fmt.Printf("no book found with ID `%s`\n", bookID)
		os.Exit(1)
	}
	fmt.Println(books.BookToString(book))
}

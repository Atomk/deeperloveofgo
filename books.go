package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	ID     string // ISBN
	Title  string
	Author string
	Copies int
}

// catalog associates a book's ID with the corresponding Book instance.
//
// the type reads "map from string to Book"
var catalog = map[string]Book{
	"9788804497882": {
		ID:     "9788804497882",
		Title:  "Artemis Fowl",
		Author: "Eoin Colfer",
		Copies: 3_000_000,
	},
	"8879262920": {
		ID:     "8879262920",
		Title:  "Bolle di sapone",
		Author: "Gianni Rodari",
		Copies: 100_000,
	},
}

/*
// Old function for reference
func printBook(book Book) {
	fmt.Println(book.Title, "by", book.Author, "sold", book.Copies, "copies")
	fmt.Printf("%v by %v sold %v copies\n", book.Title, book.Author, book.Copies)
}
*/

// Testable version of `printBook`.
func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (%v copies sold)", book.Title, book.Author, book.Copies)
}

func GetAllBooks() map[string]Book {
	return catalog
}

func GetAllBooksAsSlice() []Book {
	// The long way 1
	// booksSlice := make([]Book, len(catalog))
	// i := 0
	// for _, book := range catalog {
	// 	booksSlice[i] = book
	// 	i++
	// }
	// return booksSlice

	// The long way 2
	// booksSlice := []Book{}
	// for _, book := range catalog {
	// 	booksSlice = append(booksSlice, book)
	// }
	// return booksSlice

	// The short way (one-liner)
	// return slices.Collect(maps.Values(catalog))

	// The short way (with explanation)
	// An iterator is a sort of "lazy slice", it produces items only when requested.
	// Useful when you have a lot of elements, allows to work on one at a time
	// without having to read all elements in advance.
	iterator := maps.Values(catalog)
	// Consume the iterator and collect all elements into a slice
	return slices.Collect(iterator)
}

// Uses the “comma, ok” pattern.
func GetBook(ID string) (Book, bool) {
	// Long way
	// This works on both slices and maps - they are similar in how they are used,
	// you can think of a slice as a special case of maps where keys are consecutive
	// non-negative integers. They are implemented quite differently though.
	// `range` gives index+item for slices, key+item for maps
	// for _, book := range catalog {
	// 	if book.ID == ID {
	// 		return book, true
	// 	}
	// }
	// return Book{}, false

	// Short way
	// book := catalog[ID] // if ID is missing, book will be a zero value
	book, ok := catalog[ID]
	return book, ok
}

func AddBook(book Book) error {
	_, ok := catalog[book.ID]
	if ok {
		return fmt.Errorf("a book with ID '%s' already exists in the catalog", book.ID)
	}

	// TODO prevent adding an invalid book (empty ID or other fields)

	catalog[book.ID] = book
	return nil
}

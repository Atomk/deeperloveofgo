// TODO: why not 'books'? To avoid including this into the executable?
package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() map[string]books.Book {
	return map[string]books.Book{
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
}

// "BookToString formats book info as string"
func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "Artemis Fowl",
		Author: "Eoin Colfer",
		Copies: 3_000_000,
	}
	want := "Artemis Fowl by Eoin Colfer (3000000 copies sold)"
	got := books.BookToString(input)
	if want != got {
		// %q prints a string, inside double quotes as it's in the code
		// (i.e. it shows escape sequences)
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooksAsSlice_ReturnsCatalogBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := []books.Book{
		{
			ID:     "9788804497882",
			Title:  "Artemis Fowl",
			Author: "Eoin Colfer",
			Copies: 3_000_000,
		},
		{
			ID:     "8879262920",
			Title:  "Bolle di sapone",
			Author: "Gianni Rodari",
			Copies: 100_000,
		},
	}

	got := books.GetAllBooksAsSlice(catalog)
	// A slice is an ordered sequence, while the iteration order of maps is not
	// deterministic, as demonstrated by running this test multiple times without
	// sorting the result:
	//     go test -count=100   # run each test 100 times
	// (sometimes the test will succeed, sometimes it will fail)
	// NOTE: sorting is a generally expensive operation so it's better to sort
	// here just for the test, instead of making the function above return an
	// ordered slice. After all other then in this test we don't care about the
	// order of books.

	// Sort the slice by book author, in-place
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	// if want != got // invalid operation: want != got (slice can only be compared to nil)
	if !slices.Equal(want, got) {
		// %#v prints as Go values, like __repr__ in Python.
		// This is essential when something returns an empty string or struct,
		// which would otherwise show up as empty spaces.
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindsTheBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := books.Book{
		ID:     "8879262920",
		Title:  "Bolle di sapone",
		Author: "Gianni Rodari",
		Copies: 100_000,
	}
	got, ok := books.GetBook(catalog, want.ID)
	if !ok {
		t.Fatalf("book not found")
	}
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_DoesNotFindTheBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := books.GetBook(catalog, "fake-isbn-1234")
	if ok {
		t.Fatalf("want false (book not found), got true")
	}
}

func TestAddBook_AddsNewBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	newBook := books.Book{
		ID:     "fake1234",
		Title:  "A fake book",
		Author: "Somebody",
		Copies: 1,
	}

	// precondition
	// if this was missing, and the book already existed, the risk is that the
	// test by itself would not catch a bug (assuming the function only satisfies
	// the requirement for this test, that is just adding a book with no additional
	// behavior or check)
	// This condition ensures that AddBook actually changes the state, the postcondition
	// only ensures the state is what we expect but what caused it to get to that state
	_, ok := books.GetBook(catalog, newBook.ID)
	if ok {
		t.Fatalf("book with ID '%s' already exists", newBook.ID)
	}

	books.AddBook(catalog, newBook)

	// postcondition
	_, ok = books.GetBook(catalog, newBook.ID)
	if !ok {
		t.Fatalf("book with ID '%s' bot found, expected to find it", newBook.ID)
	}
}

func TestAddBook_CannotReplaceExistingBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	newBook := books.Book{
		ID:     "fake1234",
		Title:  "A fake book",
		Author: "Somebody",
		Copies: 1,
	}

	err := books.AddBook(catalog, newBook)
	if err != nil {
		t.Fatal("cannot insert the test book")
	}

	err = books.AddBook(catalog, newBook)
	if err == nil {
		t.Fatal("want an error, got nil")
	}
}

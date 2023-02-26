package main

import (
	"fmt"
	"time"
)

type Title string // book title
type Name string  // library member name

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total int // total owned by the library
	lent  int // total books lent out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) { //
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	for bookTitle, book := range library.books {
		fmt.Println(bookTitle, " [available: ", book.total-book.lent, "/", book.total, "]")
	}
	fmt.Println()
}

func findBookInLibrary(library *Library, title Title) (BookEntry, bool) {
	book, found := library.books[title]
	if !found {
		fmt.Println("[", title, "] Book not part of the library!")
	}
	return book, found
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := findBookInLibrary(library, title)
	if !found {
		return false
	}

	if book.lent == book.total {
		fmt.Println("[", title, "] All books landed out!")
		return false
	}
	book.lent++
	library.books[title] = book

	member.books[title] = LendAudit{
		checkOut: time.Now(),
	}

	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := findBookInLibrary(library, title)
	if !found {
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("[", title, "] Book not lent by member ", member.name)
		return false
	}

	book.lent--
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{total: 4, lent: 0}
	library.books["Go to the bank"] = BookEntry{total: 3, lent: 0}
	library.books["Lets go to the maaaaaall"] = BookEntry{total: 2, lent: 0}
	library.books["GooGoo Dolls"] = BookEntry{total: 1, lent: 0}
	printLibraryBooks(&library)

	library.members["Medo"] = Member{"Medo", make(map[Title]LendAudit)}
	library.members["Katya"] = Member{"Katya", make(map[Title]LendAudit)}
	library.members["Ada"] = Member{"Ada", make(map[Title]LendAudit)}
	printMemberAudits(&library)

	member := library.members["Medo"]
	checkedOut := checkoutBook(&library, "Lets go to the maaaaaall", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Lets go to the maaaaaall", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

}

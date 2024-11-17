package main

import (
	"fmt"
)

type Book struct {
	Title string
	Read  bool
}

func checkAllBooksRead(books []Book) bool {
	for _, book := range books {
		if !book.Read {
			return false
		}
	}
	return true
}

func main() {
	readingList := []Book{
		{Title: "Jungle Book", Read: true},
		{Title: "Ride or Die", Read: false},
		{Title: "Holy Ghost Fire", Read: true},
		{Title: "The Breeder", Read: false},
	}

	allRead := checkAllBooksRead(readingList)

	if allRead {
		fmt.Println("Congratulations! You have read all the books on your list.")
	} else {
		fmt.Println("You still have books to read!")
	}
}

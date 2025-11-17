package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Isbn   string
	Weight int
}

func main() {

	books := []Book{
		{
			Title:  "Moby Dick",
			Isbn:   "123-567",
			Weight: 125,
		},
		{
			Title:  "Dick Turpin",
			Isbn:   "567-432",
			Weight: 200,
		},
	}

	for _, book := range books {
		fmt.Println(book)
	}
}

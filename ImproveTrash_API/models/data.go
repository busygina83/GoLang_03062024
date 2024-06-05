package models

import "slices"

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID:            1,
		Title:         "Lord of the Rings. Vol.1",
		YearPublished: 1978,
		Author: Author{
			Name:     "J.R",
			LastName: "Tolkin",
			BornYear: 1892,
		},
	}
	DB = append(DB, book1)
}

func FindBookById(id int) (Book, bool) {
	var book Book
	var found bool
	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
			break
		}
	}
	return book, found
}

func DeleteBookById(id int) bool {
	for idx, b := range DB {
		if b.ID == id {
			// DB = append(DB[:idx], DB[idx + 1:]...)
			DB = slices.Delete(DB, idx, idx+1)
			return true
		}
	}

	return false
}

func UpdateBookById(id int, book Book) bool {
	found := false
	for k, b := range DB {
		if b.ID == id {
			if len(book.Title) > 0 {
				DB[k].Title = book.Title
			}

			if len(book.Author.Name) > 0 {
				DB[k].Author.Name = book.Author.Name
			}

			if len(book.Author.LastName) > 0 {
				DB[k].Author.LastName = book.Author.LastName
			}

			if book.Author.BornYear > 0 {
				DB[k].Author.BornYear = book.Author.BornYear
			}

			if book.YearPublished > 0 {
				DB[k].YearPublished = book.YearPublished
			}
			found = true
			
		}
	}
	return found
}
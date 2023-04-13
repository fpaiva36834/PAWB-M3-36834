package service

import (
	"awesomeProject/entity"
	"errors"
)

var Books []entity.Book

func InitializeBooks() {
	Books = []entity.Book{}
}

func GetAllBooks() []entity.Book {
	return Books
}

func InsertBook(book entity.Book) entity.Book {
	book.ID = uint64(len(Books) + 1)
	Books = append(Books, book)
	return book
}

func GetBookByID(id uint64) (entity.Book, error) {
	for i := range Books {
		if Books[i].ID == id {
			return Books[i], nil
		}
	}
	return entity.Book{}, errors.New("book not found")
}

func UpdateBookByID(book entity.Book) (entity.Book, error) {
	for i := range Books {
		if Books[i].ID == book.ID {
			Books[i] = book
			return Books[i], nil
		}
	}
	return entity.Book{}, errors.New("book not found")
}

func DeleteBookByID(id uint64) error {
	for i := range Books {
		if Books[i].ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

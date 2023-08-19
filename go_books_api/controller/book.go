package controller

import (
	"github.com/faaizz/code_demos/go_books_api/model"
)

type Book struct{}

func (b *Book) GetBooks() ([]model.Book, error) {
	var books []model.Book
	tx := DB.Find(&books)
	return books, tx.Error
}

func (b *Book) CreateBook(title, author string, year int) (model.Book, error) {
	book := model.Book{
		Title:  title,
		Author: author,
		Year:   year,
	}

	tx := DB.Create(&book)
	return book, tx.Error
}

func (b *Book) ReadBook(id uint) (model.Book, error) {
	var book model.Book
	tx := DB.First(&book, id)
	return book, tx.Error
}

func (b *Book) UpdateBook(id uint, title, author string, year int) (model.Book, error) {
	var book model.Book
	tx := DB.First(&book, id)
	if tx.Error != nil {
		return book, tx.Error
	}

	book.Title = title
	book.Author = author
	book.Year = year
	tx = DB.Save(&book)

	return book, tx.Error
}

func (b *Book) DeleteBook(id uint) error {
	tx := DB.Delete(&model.Book{}, id)
	return tx.Error
}

package controller_test

import (
	"testing"

	"github.com/faaizz/code_demos/go_books_api/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type booksTestCase struct {
	description string
	hasError    bool
	err         error
}

var testCases []booksTestCase

func init() {
	controller.BC = &controller.Book{}
	testCases = []booksTestCase{
		{
			"should return no error",
			false,
			nil,
		},
		{
			"should return record not found error",
			true,
			gorm.ErrRecordNotFound,
		},
		{
			"should return where clause error",
			true,
			gorm.ErrMissingWhereClause,
		},
	}
}

func TestGetBooks(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			dummyDB := &gorm.DB{}
			dummyDB.Error = tc.err

			mockDB := NewMockIGormDB(t)
			mockDB.On("Find", mock.Anything).Return(dummyDB)
			controller.DB = mockDB
			// Act
			_, err := controller.BC.GetBooks()
			// Assert
			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			dummyDB := &gorm.DB{}
			dummyDB.Error = tc.err

			mockDB := NewMockIGormDB(t)
			mockDB.On("Create", mock.Anything).Return(dummyDB)
			controller.DB = mockDB
			// Act
			_, err := controller.BC.CreateBook("", "", 0)
			// Assert
			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestReadBook(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			dummyDB := &gorm.DB{}
			dummyDB.Error = tc.err

			mockDB := NewMockIGormDB(t)
			mockDB.On("First", mock.Anything, mock.Anything).Return(dummyDB)
			controller.DB = mockDB
			// Act
			_, err := controller.BC.ReadBook(0)
			// Assert
			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	type updateBooksTestCase struct {
		description   string
		hasError      bool
		firstHasError bool
		err           error
	}

	testCases := []updateBooksTestCase{
		{
			"should return no error",
			false,
			false,
			nil,
		},
		{
			"should return record not found error",
			true,
			true,
			gorm.ErrRecordNotFound,
		},
		{
			"should return duplicated key error",
			true,
			false,
			gorm.ErrDuplicatedKey,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			dummyDB := &gorm.DB{}
			dummyDB.Error = tc.err

			mockDB := NewMockIGormDB(t)

			firstDB := &gorm.DB{}
			if tc.firstHasError {
				firstDB.Error = tc.err
			}
			mockDB.On("First", mock.Anything, mock.Anything).Return(firstDB)
			if !tc.firstHasError {
				mockDB.On("Save", mock.Anything).Return(dummyDB)
			}
			controller.DB = mockDB
			// Act
			_, err := controller.BC.UpdateBook(0, "", "", 0)
			// Assert
			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			dummyDB := &gorm.DB{}
			dummyDB.Error = tc.err

			mockDB := NewMockIGormDB(t)
			mockDB.On("Delete", mock.Anything, mock.Anything).Return(dummyDB)
			controller.DB = mockDB
			// Act
			err := controller.BC.DeleteBook(0)
			// Assert
			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

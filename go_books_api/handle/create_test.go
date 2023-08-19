package handle_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/faaizz/code_demos/go_books_api/controller"
	"github.com/faaizz/code_demos/go_books_api/handle"
	"github.com/faaizz/code_demos/go_books_api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockResponseWriter struct {
	mock.Mock
	response string
}

func (m *MockResponseWriter) Header() http.Header {
	return http.Header{}
}
func (m *MockResponseWriter) Write(bs []byte) (int, error) {
	m.response = string(bs[:])
	return 0, nil
}
func (m *MockResponseWriter) WriteHeader(int) {}

func TestBookCreate(t *testing.T) {
	testCases := []struct {
		description         string
		title               string
		author              string
		year                int
		hasRequestBodyError bool
		hasControllerError  bool
		errStr              string
	}{
		{
			"should return no error",
			"test",
			"test",
			2020,
			false,
			false,
			"",
		},
		{
			"should return could not decode request body error",
			"test",
			"test",
			2020,
			true,
			false,
			"could not decode request body",
		},
		{
			"should return could not create book error",
			"test",
			"test",
			2020,
			false,
			true,
			"could not create book",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			book := model.Book{
				Model:  gorm.Model{},
				Title:  tc.title,
				Author: tc.author,
				Year:   tc.year,
			}

			rBodyStr := fmt.Sprintf(`{"title":"%s","author":"%s","year":%d}`, tc.title, tc.author, tc.year)
			if tc.hasRequestBodyError {
				rBodyStr = fmt.Sprintf(`"it__e":"%s","author":"%s","year":%d}`, tc.title, tc.author, tc.year)
			}

			mockBC := NewMockIController(t)
			var mockBCErr interface{}
			if tc.hasControllerError {
				mockBCErr = fmt.Errorf("could not create book")
			} else {
				mockBCErr = nil
			}
			if !tc.hasRequestBodyError {
				mockBC.On("CreateBook", tc.title, tc.author, tc.year).Return(book, mockBCErr)
			}
			controller.BC = mockBC

			httpReq := http.Request{
				Body: io.NopCloser(strings.NewReader(rBodyStr)),
			}
			httpRes := MockResponseWriter{}

			// Act
			handle.BookCreate(&httpRes, &httpReq, nil)

			// Assert
			if tc.hasControllerError || tc.hasRequestBodyError {
				assert.Equal(t, httpRes.response, tc.errStr+"\n")
			}
		})
	}
}

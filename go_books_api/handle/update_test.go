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
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBookUpdate(t *testing.T) {
	testCases := []struct {
		description         string
		idStr               string
		id                  uint
		hasSanitizeIDError  bool
		title               string
		author              string
		year                int
		hasRequestBodyError bool
		hasControllerError  bool
		errStr              string
	}{
		{
			"should return no error",
			"0",
			0,
			false,
			"test",
			"test",
			2020,
			false,
			false,
			"",
		},
		{
			"should return could not decode request body error",
			"0",
			0,
			false,
			"test",
			"test",
			2020,
			true,
			false,
			"could not decode request body",
		},
		{
			"should return could not update book error",
			"0",
			0,
			false,
			"test",
			"test",
			2020,
			false,
			true,
			"failed to update book",
		},
		{
			"should return id must be valid integer error",
			"poo",
			0,
			true,
			"test",
			"test",
			2020,
			false,
			false,
			"id must be a valid integer",
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
				mockBCErr = fmt.Errorf("failed to update book")
			} else {
				mockBCErr = nil
			}
			if !tc.hasRequestBodyError && !tc.hasSanitizeIDError {
				mockBC.On("UpdateBook", tc.id, tc.title, tc.author, tc.year).Return(book, mockBCErr)
			}
			controller.BC = mockBC

			httpReq := http.Request{
				Body: io.NopCloser(strings.NewReader(rBodyStr)),
			}
			httpRes := MockResponseWriter{}
			ps := httprouter.Params{
				httprouter.Param{
					Key:   "id",
					Value: tc.idStr,
				},
			}

			// Act
			handle.BookUpdate(&httpRes, &httpReq, ps)

			// Assert
			if tc.hasControllerError || tc.hasRequestBodyError || tc.hasSanitizeIDError {
				assert.Equal(t, httpRes.response, tc.errStr+"\n")
			}
		})
	}
}

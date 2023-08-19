package handle_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/faaizz/code_demos/go_books_api/controller"
	"github.com/faaizz/code_demos/go_books_api/handle"
	"github.com/faaizz/code_demos/go_books_api/model"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestBookRead(t *testing.T) {
	testCases := []struct {
		description        string
		idStr              string
		id                 uint
		hasSanitizeIDError bool
		hasControllerError bool
		errStr             string
	}{
		{
			"should return no error",
			"0",
			0,
			false,
			false,
			"",
		},
		{
			"should return could not read book error",
			"0",
			0,
			false,
			true,
			"could not read book",
		},
		{
			"should return id is required error",
			"",
			0,
			true,
			false,
			"id is required",
		},
		{
			"should return invalid id error",
			"poo",
			0,
			true,
			false,
			"id must be a valid integer",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			mockBC := NewMockIController(t)
			var mockBCErr interface{}
			if tc.hasControllerError {
				mockBCErr = fmt.Errorf("could not read book")
			} else {
				mockBCErr = nil
			}
			if !tc.hasSanitizeIDError {
				mockBC.On("ReadBook", tc.id).Return(model.Book{}, mockBCErr)
			}
			controller.BC = mockBC

			httpReq := http.Request{}
			httpRes := MockResponseWriter{}
			ps := httprouter.Params{
				httprouter.Param{
					Key:   "id",
					Value: tc.idStr,
				},
			}

			// Act
			handle.BookRead(&httpRes, &httpReq, ps)

			// Assert
			if tc.hasControllerError || tc.hasSanitizeIDError {
				assert.Equal(t, httpRes.response, tc.errStr+"\n")
			}
		})
	}
}

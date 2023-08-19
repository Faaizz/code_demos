package handle_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/faaizz/code_demos/go_books_api/controller"
	"github.com/faaizz/code_demos/go_books_api/handle"
	"github.com/faaizz/code_demos/go_books_api/model"
	"github.com/stretchr/testify/assert"
)

func TestBookIndex(t *testing.T) {
	testCases := []struct {
		description        string
		hasControllerError bool
		errStr             string
	}{
		{
			"should return no error",
			false,
			"",
		},
		{
			"should return could not get books error",
			true,
			"could not get books",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			mockBC := NewMockIController(t)
			var mockBCErr interface{}
			if tc.hasControllerError {
				mockBCErr = fmt.Errorf("could not get books")
			} else {
				mockBCErr = nil
			}
			mockBC.On("GetBooks").Return([]model.Book{}, mockBCErr)
			controller.BC = mockBC

			httpReq := http.Request{}
			httpRes := MockResponseWriter{}

			// Act
			handle.BookIndex(&httpRes, &httpReq, nil)

			// Assert
			if tc.hasControllerError {
				assert.Equal(t, httpRes.response, tc.errStr+"\n")
			}
		})
	}
}

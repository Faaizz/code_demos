package handle

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/faaizz/code_demos/go_books_api/controller"
	"github.com/faaizz/code_demos/go_books_api/model"
)

func BookCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	br := &model.BookRequest{}

	err := json.NewDecoder(r.Body).Decode(br)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not decode request body", http.StatusBadRequest)
		return
	}

	err = br.Validate()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := controller.BC.CreateBook(br.Title, br.Author, br.Year)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not create book", http.StatusInternalServerError)
		return
	}

	w = addHeaders(w)
	json.NewEncoder(w).Encode(book)
}

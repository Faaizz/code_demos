package handle

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/faaizz/code_demos/go_books_api/controller"
)

func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bs, err := controller.BC.GetBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, "could not get books", http.StatusInternalServerError)
		return
	}

	w = addHeaders(w)
	json.NewEncoder(w).Encode(bs)
}

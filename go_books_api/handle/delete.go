package handle

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/faaizz/code_demos/go_books_api/controller"
)

func BookDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := sanitizeID(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.BC.DeleteBook(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not delete book", http.StatusInternalServerError)
		return
	}

	w = addHeaders(w)
	fmt.Fprint(w, "{}")
}

package handle

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func ApiDocs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/yaml; charset=utf-8")
	data, err := os.ReadFile("api-docs.yml")
	if err != nil {
		log.Println(err)
		http.Error(w, "could not read api docs", http.StatusNotFound)
		return
	}
	w = addHeaders(w)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(data))
}

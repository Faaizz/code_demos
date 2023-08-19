package handle

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Status struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := Status{"OK"}

	w.WriteHeader(http.StatusOK)
	w = addHeaders(w)
	json.NewEncoder(w).Encode(s)
}

package middleware

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var requiredUser string
var requiredPassword string

func init() {
	requiredUser = os.Getenv("BASIC_AUTH_USER")
	if requiredUser == "" {
		panic("BASIC_AUTH_USER environment variable required but not set")
	}
	requiredPassword = os.Getenv("BASIC_AUTH_PASSWORD")
	if requiredPassword == "" {
		panic("BASIC_AUTH_PASSWORD environment variable required but not set")
	}
}

func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

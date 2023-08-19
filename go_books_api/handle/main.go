package handle

import (
	"errors"
	"net/http"
	"strconv"
)

func addHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
}

func sanitizeID(idStr string) (uint, error) {
	if idStr == "" {
		return 0, errors.New("id is required")
	}
	idUint64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, errors.New("id must be a valid integer")
	}
	return uint(idUint64), nil
}

package web

import (
	"fmt"
	"net/http"
)

func plain(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, http.StatusText(http.StatusOK))
}

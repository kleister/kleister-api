package respond

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// JSON is a simple helper to render JSON responses.
func JSON(w http.ResponseWriter, _ *http.Request, v any) {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(v)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to generate json response")

		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}

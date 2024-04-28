package get

import (
	"client/internal/handlers"
	"client/internal/lib/log/sl"
	"client/internal/views/pages"
	"log/slog"
	"net/http"
)

// IndexHandler handles the GET / route.
func IndexHandler(log *slog.Logger) http.HandlerFunc {
	const op = "get.IndexHandler"
	log = log.With(slog.String("op", op))

	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("IndexHandler called")

		c := pages.IndexPage("Jake")

		err := handlers.Render(w, r, c)
		if err != nil {
			log.Error("Failed to render page", sl.Err(err))
		}
	}
}

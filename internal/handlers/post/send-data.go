package post

import (
	"client/internal/handlers"
	"client/internal/lib/log/sl"
	"client/internal/views/components"
	"log/slog"
	"net/http"
)

// SendDataHandler handles the POST /send-data route.
func SendDataHandler(log *slog.Logger) http.HandlerFunc {
	const op = "post.SendDataHandler"
	log = log.With(slog.String("op", op))

	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("SendDataHandler called")

		data := r.FormValue("data")

		// Validate data (example)
		if len(data) < 3 {
			log.Debug("Invalid data received")

			w.WriteHeader(http.StatusUnauthorized)
			c := components.ErrorMessage("Data is too short")
			err := handlers.Render(w, r, c)
			if err != nil {
				log.Error("Failed to render error message", sl.Err(err))
			}
			return
		}

		log.Debug("Data received", slog.String("data", data))

		w.WriteHeader(http.StatusOK)
		c := components.Success(data)
		err := handlers.Render(w, r, c)
		if err != nil {
			log.Error("Failed to render success message", sl.Err(err))
		}
	}
}

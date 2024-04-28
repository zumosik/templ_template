package main

import (
	"client/internal/config"
	"client/internal/handlers/get"
	"client/internal/handlers/post"
	"client/internal/lib/log/sl"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// get config
	cfg := config.MustLoadConfig()
	// setup logger
	log := sl.MustConfigureLogger(cfg.Env)

	// setup router
	router := chi.NewMux()

	// serve static files
	fileServer := http.FileServer(http.Dir("./public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	// Get
	router.Get("/", get.IndexHandler(log))

	// Post
	router.Post("/send-data", post.SendDataHandler(log))

	// Start server
	log.Info("Starting server", slog.String("addr", fmt.Sprintf(":%d", cfg.Port)))
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	if err != nil {
		log.Error("Failed to start server", sl.Err(err))
	}
}

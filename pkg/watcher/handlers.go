package watcher

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (w *Watcher) getHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthcheck"))

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Gobble 🦃")) // goli
	})

	r.Get("/users/", getUsersHandler(w))

	return r
}

func getUsersHandler(w *Watcher) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		b, err := json.MarshalIndent(w.Users, "", "  ")

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)
	}
}

package watcher

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gobble/pkg/tasks"
	"net/http"
)

func (w *Watcher) getHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthcheck"))

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Gobble 🦃"))
	})

	r.Get("/users/update", updateUsers(w))

	return r
}

func updateUsers(w *Watcher) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		u, err := tasks.GetUsers(w.Config)

		if err != nil {
			writer.Write([]byte(err.Error()))
		}
		e := json.NewEncoder(writer)
		e.Encode(u)
	}
}

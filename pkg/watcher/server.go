package watcher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"syscall"
	"time"
)

// startServer starts the http server. If the ListenAndServe method returns an error, and the error
// is not ErrServerClosed, it will send an interrupt signal to the watcher to stop all services
func (w *Watcher) startServer() error {
	addr := fmt.Sprintf("%s:%s", w.Config.Host, w.Config.Port)

	w.server = &http.Server{
		Handler: w.getHandler(),
		Addr:    addr,
	}

	log.Printf("starting server on address %q\n", addr)

	if err := w.server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
		w.sigkill <- syscall.SIGINT
	}

	return nil

}

// stopServer stops the http server by calling server.Shutdown with the provided context
func (w *Watcher) stopServer(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	if err := w.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error occurred while shutting down server: %w", err)
	}

	return nil
}

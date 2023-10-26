package watcher

import (
	"context"
	flag "github.com/spf13/pflag"
	"gobble/pkg/configuration"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Watcher keeps track of all the things
type Watcher struct {
	Config *configuration.Config
	Flags  *configuration.Flags
	server *http.Server
	// sigkill causes immediate program termination, cannot be handled or ignored
	sigkill chan os.Signal
	// sighup signal is sent to the process when the controlling terminal is closed
	sighup chan os.Signal
}

func Run() error {
	// Get flags
	flags := &configuration.Flags{
		FlagSet:    flag.NewFlagSet("gobble", flag.ExitOnError),
		ConfigFile: os.Getenv("GOBBLE_CONFIG_FILE"),
		DryRun:     false,
	}

	if err := flags.Parse(os.Args[:1]); err != nil {
		return err
	}

	// Read config
	cfg := configuration.New()

	if err := cfg.Parse(flags); err != nil {
		return err
	}

	// Set up Watcher with some default values
	watcher := &Watcher{
		Config:  cfg,
		Flags:   flags,
		sigkill: make(chan os.Signal, 1),
		sighup:  make(chan os.Signal, 1),
		server:  &http.Server{},
	}

	ctx := context.Background()
	return watcher.start(ctx)
}

func (w *Watcher) start(ctx context.Context) error {
	log.Println("starting watcher 🦃")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go w.startServer()
	w.checkUserMappings()

	// Channel for catching ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	for {
		select {
		case sig := <-w.sigkill:
			log.Printf("received termination signal: %v", sig)
			return w.stop(ctx)
		case sig := <-c:
			log.Printf("received interrupt signal: %v", sig)
			return w.stop(ctx)
		}
	}
}

func (w *Watcher) stop(ctx context.Context) error {
	return w.stopServer(ctx)
}

func (w *Watcher) checkUserMappings() {
	for _, m := range w.Config.UserMappings {
		// TODO: actually do something with these user mappings
		log.Println(m.String())
	}
}

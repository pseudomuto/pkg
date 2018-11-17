package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

// ServerOpts describes options for running HTTP servers
type ServerOpts struct {
	// The Done channel is closed when the server has been shutdown
	Done chan interface{}
	// Logger is the FieldLogger (from logrus) used for logging
	Logger log.FieldLogger
	// Signals refers to the signals that should trigger a shutdown
	Signals []os.Signal
	// The channel that is notified when signals are received
	SignalTrap chan os.Signal
	// The duration to wait for graceful shutdown
	ShutdownTimeout time.Duration
}

// DefaultServerOpts creates a ServerOpts object with defaults applied
func DefaultServerOpts() ServerOpts {
	host, _ := os.Hostname()

	return ServerOpts{
		Done:            make(chan interface{}),
		Logger:          log.WithField("host", host),
		Signals:         []os.Signal{syscall.SIGINT, syscall.SIGTERM},
		SignalTrap:      make(chan os.Signal, 1),
		ShutdownTimeout: 5 * time.Second,
	}
}

// RunServer starts the given http.Server instance and sets up support for graceful shutdown via signal traps.
func RunServer(svr *http.Server, opts ServerOpts) {
	signal.Notify(opts.SignalTrap, opts.Signals...)

	go func() {
		sig := <-opts.SignalTrap
		opts.Logger.WithField("signal", sig).Info("received shutdown signal")

		ctx, cancel := context.WithTimeout(context.Background(), opts.ShutdownTimeout)
		defer cancel()

		// turn off keep alives for new connections
		svr.SetKeepAlivesEnabled(false)
		if err := svr.Shutdown(ctx); err != nil {
			opts.Logger.WithError(err).WithField("timeout", opts.ShutdownTimeout).Warn("failed to shutdown gracefully")
		}

		close(opts.Done)
	}()

	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			opts.Logger.WithError(err).Fatalln("failed to start/stop server cleanly")
		}
	}()
}

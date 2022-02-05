package main

import (
	"context"
	"errors"
	"fmt"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/coreos/go-systemd/v22/activation"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
)

const (
	httpIOTimeout       = 10 * time.Second
	httpShutdownTimeout = 5 * time.Second
)

func main() {
	if err := run(); err != nil {
		log.Err(err).Msg("failed to run app")
		os.Exit(1)
	}
}

func run() error {
	lg := initLog()

	ctx, stopFn := initSignals(context.Background())
	defer stopFn()

	handler := newHandler(lg)
	server := newServer(ctx, lg, "localhost:8080", handler)

	if err := runServer(ctx, server); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	lg.Warn().Msg("exiting")

	return nil
}

func initSignals(ctx context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, os.Interrupt, unix.SIGTERM)
}

func newServer(ctx context.Context, lg zerolog.Logger, addr string, handler http.Handler) *http.Server {
	// Store the logger in the context, so trickles down to requests.
	ctx = lg.WithContext(ctx)

	server := http.Server{
		Addr:    addr,
		Handler: handler,

		ReadTimeout:  httpIOTimeout,
		WriteTimeout: httpIOTimeout,

		ErrorLog: stdlog.New(log.Logger, "", 0),

		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	return &server
}

func runServer(ctx context.Context, server *http.Server) error {
	listeners, err := activation.Listeners()
	if err != nil {
		return fmt.Errorf("failed to get systemd listeners: %w", err)
	}

	if len(listeners) > 1 {
		return fmt.Errorf("unexpected number of listeners: %v", len(listeners))
	}

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		var err error

		if len(listeners) > 0 {
			log.Info().Str("addr", "systemd-socket").Msg("starting server")
			err = server.Serve(listeners[0])
		} else {
			log.Info().Str("addr", server.Addr).Msg("starting server")
			err = server.ListenAndServe()
		}

		if err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server exited unexpectedly: %w", err)
		}

	case <-ctx.Done():
		log.Warn().Msg("signal received")

		// Recevied a stop signal, give the server a chance to close connections
		// cleanly.
		ctxShutdown, cancel := context.WithTimeout(context.Background(), httpShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctxShutdown); err != nil {
			return fmt.Errorf("failed to cleanly shut down server: %w", err)
		}
	}

	return nil
}

// newListener returns a new listening socket, if the server is started with
// systemd socket activation, nil otherwise.
func newListener() (net.Listener, error) {
	nfds, err := strconv.Atoi(os.Getenv("LISTEN_FDS"))
	if err != nil || nfds == 0 {
		return nil, nil
	}

	l, err := net.FileListener(os.NewFile(3, "systemd socket-activation fd"))
	if err != nil {
		return nil, fmt.Errorf("failed to create file listener: %w", err)
	}

	return l, nil
}

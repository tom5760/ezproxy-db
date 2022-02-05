package main

import (
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/journald"
	"github.com/rs/zerolog/log"
)

func initLog() zerolog.Logger {
	consolew := zerolog.NewConsoleWriter()
	consolew.Out = os.Stderr

	journalw := journald.NewJournalDWriter()

	writer := zerolog.MultiLevelWriter(consolew, journalw)

	lg := zerolog.New(writer).With().
		Timestamp().
		Logger()

	log.Logger = lg

	return lg
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rv := recover(); rv != nil && rv != http.ErrAbortHandler {
				evt := hlog.FromRequest(r).Error().
					Bytes(zerolog.ErrorStackFieldName, debug.Stack())

				if err, ok := rv.(error); ok {
					evt = evt.Err(err)
				} else {
					evt = evt.Interface("err", rv)
				}

				evt.Msg("handler panicked")

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// logAccessFn logs details about the results of a request to zerolog.  It is
// used with hlog.AccessHandler.
func logAccessFn(r *http.Request, status, size int, duration time.Duration) {
	level := zerolog.ErrorLevel

	if status >= 200 && status <= 399 {
		level = zerolog.InfoLevel
	}

	hlog.FromRequest(r).WithLevel(level).
		Int("status", status).
		Int("size", size).
		Dur("duration", duration).
		Msg("request handled")
}

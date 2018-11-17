package http

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// WithLogging wraps the supplied handler to handle request logging
func WithLogging(logger log.FieldLogger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(w, r)

		duration := time.Since(startTime)
		logger.WithFields(log.Fields{
			"dur":    duration,
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("processed request")
	})
}

package http

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// WithLogging wraps the supplied handler to handle request logging
func WithLogging(logger log.FieldLogger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		sw := &statusResponseWriter{w, http.StatusOK}
		h.ServeHTTP(sw, r)

		duration := time.Since(startTime)
		logger.WithFields(log.Fields{
			"dur":    duration,
			"method": r.Method,
			"path":   r.URL.Path,
			"status": sw.status,
		}).Info("processed request")
	})
}

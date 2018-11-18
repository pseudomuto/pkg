package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/pseudomuto/pkg/http"
	logTest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestWithLogging(t *testing.T) {
	tests := []struct {
		path   string
		method string
		status int
	}{
		{path: "/some/path", method: "GET", status: 200},
		{path: "/thing/create", method: "POST", status: 201},
		{path: "/thing/unknown", method: "DELETE", status: 404},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(test.method, test.path, nil)
		logger, hook := logTest.NewNullLogger()

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(test.status)
		})

		WithLogging(logger, h).ServeHTTP(w, r)

		entry := hook.LastEntry()
		assert.Equal(t, test.method, entry.Data["method"])
		assert.Equal(t, test.path, entry.Data["path"])
		assert.Equal(t, test.status, entry.Data["status"])
	}
}

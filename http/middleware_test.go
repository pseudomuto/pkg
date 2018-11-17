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
	}{
		{path: "/some/path", method: "GET"},
		{path: "/thing/create", method: "POST"},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(test.method, test.path, nil)
		logger, hook := logTest.NewNullLogger()

		h := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
		WithLogging(logger, h).ServeHTTP(w, r)

		entry := hook.LastEntry()
		assert.Equal(t, test.method, entry.Data["method"])
		assert.Equal(t, test.path, entry.Data["path"])
	}
}

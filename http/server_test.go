package http_test

import (
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	. "github.com/pseudomuto/pkg/http"
	"github.com/stretchr/testify/assert"
)

func TestDefaultServerOpts(t *testing.T) {
	opts := DefaultServerOpts()

	assert.NotNil(t, opts.Done)
	assert.NotNil(t, opts.Logger)
	assert.Equal(t, []os.Signal{syscall.SIGINT, syscall.SIGTERM}, opts.Signals)
	assert.NotNil(t, opts.SignalTrap)
	assert.Equal(t, 5*time.Second, opts.ShutdownTimeout)
}

func TestRunServer(t *testing.T) {
	svr := &http.Server{
		Addr:    ":5101",
		Handler: http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}),
	}

	opts := DefaultServerOpts()
	RunServer(svr, opts)

	res, err := http.Get("http://127.0.0.1:5101/path")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	opts.SignalTrap <- syscall.SIGINT
	<-opts.Done
}

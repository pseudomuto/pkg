package exec_test

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"testing"

	. "github.com/pseudomuto/pkg/exec"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	require.Equal(t, ErrEmptyCommand, Run([]string{}))

	tests := []struct {
		cmd     []string
		in      string
		out     string
		err     string
		exitErr string
	}{
		{
			cmd: []string{"echo", "-n", "test"},
			out: "test",
		},
		{
			cmd: []string{"tr", "[:lower:]", "[:upper:]"},
			in:  "test",
			out: "TEST",
		},
		{
			cmd:     []string{"which", "jbjhvyvkv"},
			exitErr: "exit status 1",
		},
	}

	for _, tt := range tests {
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		opts := []Option{Out(stdout), Err(stderr)}

		if tt.in != "" {
			opts = append(opts, In(strings.NewReader(tt.in)))
		}

		err := Run(tt.cmd, opts...)
		if tt.exitErr != "" {
			err, ok := err.(*exec.ExitError)
			require.True(t, ok)
			require.EqualError(t, err, tt.exitErr)
		} else {
			require.NoError(t, err)
		}

		require.Equal(t, tt.out, stdout.String())
		require.Equal(t, tt.err, stderr.String())
	}
}

func TestRunContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // pre-cancel

	require.EqualError(t, RunContext(ctx, []string{"which", "go"}), "context canceled")
}

func TestPipe(t *testing.T) {
	tests := []struct {
		cmds [][]string
		in   string
		out  string
		err  string
	}{
		{
			cmds: [][]string{
				{"echo", "-n", "test"},
				{"tr", "[:lower:]", "[:upper:]"},
			},
			out: "TEST",
		},
		{
			cmds: [][]string{
				{"tr", "[:lower:]", "[:upper:]"},
			},
			in:  "tEst",
			out: "TEST",
		},
	}

	for _, tt := range tests {
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		opts := []Option{Out(stdout), Err(stderr)}
		if tt.in != "" {
			opts = append(opts, In(strings.NewReader(tt.in)))
		}

		require.NoError(t, Pipe(tt.cmds, opts...))
		require.Equal(t, tt.out, stdout.String())
		require.Equal(t, tt.err, stderr.String())
	}
}

func TestPipeContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // pre-cancel

	require.EqualError(t, PipeContext(ctx, [][]string{{"which", "go"}}), "context canceled")
}

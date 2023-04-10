package exec

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
)

var (
	// ErrEmptyCommand is returned when no commands are supplied.
	ErrEmptyCommand = fmt.Errorf("no command given")
)

// Options defines in/out/err streams for commands.
type Options struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// Run runs the command using the supplied options.
func Run(cmd []string, options ...Option) error {
	return RunContext(context.Background(), cmd, options...)
}

// RunContext runs the command using the given context and supplied options.
func RunContext(ctx context.Context, cmd []string, options ...Option) error {
	if len(cmd) == 0 {
		return ErrEmptyCommand
	}

	p := cmd[0]
	args := []string{}
	if len(cmd) > 1 {
		args = cmd[1:]
	}

	opts := newOptions(options)
	execCmd := exec.CommandContext(ctx, p, args...)
	execCmd.Stdout = opts.out()
	execCmd.Stderr = opts.err()
	execCmd.Stdin = opts.in()

	return execCmd.Run()
}

// Pipe executes the given commands piping the output from the previous command
// as the input to the next.
func Pipe(cmds [][]string, options ...Option) error {
	return PipeContext(context.Background(), cmds, options...)
}

// PipeContext executes the given commands piping the output from the previous
// command as the input to the next.
func PipeContext(ctx context.Context, cmds [][]string, options ...Option) error {
	var out io.Writer = new(bytes.Buffer)
	var err io.Writer = new(bytes.Buffer)

	opts := newOptions(options)
	for i, cmd := range cmds {
		in := opts.in()
		if i != 0 {
			// Use the result of the last command as input to the next one.
			in = bytes.NewBuffer(out.(*bytes.Buffer).Bytes())
		}

		out = new(bytes.Buffer)
		err = new(bytes.Buffer)

		// Last command, so capture out/err to the original streams.
		if i == len(cmds)-1 {
			out = opts.out()
			err = opts.err()
		}

		if err := RunContext(ctx, cmd, In(in), Out(out), Err(err)); err != nil {
			return err
		}
	}

	return nil
}

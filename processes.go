package pwn

import (
	"io"
	"os"
	"os/exec"
	"time"
)

// Start starts cmd and returns a Process for it
func Start(cmd *exec.Cmd) (Process, error) {
	// file descriptors
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return Process{}, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return Process{}, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return Process{}, err
	}

	err = cmd.Start()
	if err != nil {
		return Process{}, err
	}
	return Process{
		cmd:    cmd,
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		// the maximum line length to be used with ReadLine
		maxLen: MaxLenDefault,
	}, nil
}

// Spawn spawns a new process and returns it
func Spawn(path string, args ...string) (Process, error) {
	cmd := exec.Command(path, args...)
	return Start(cmd)
}

// Process represents a spawned process
// It has the methods of a os.Process and os.Cmd
type Process struct {
	// the underlying cmd
	cmd *exec.Cmd

	// file descriptors we can manipulate
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Stderr io.ReadCloser

	// the max length to be used with ReadLine
	maxLen int
}

// WriteLine writes a line to the standard input of the running process
// t can be anything convertable to []byte (see ToBytes function)
// ToBytes will panic if it fails to convert to bytes
func (p Process) WriteLine(t interface{}) error {
	// write the data to the processes standard input
	return WriteLine(p.Stdin, t)
}

// ReadLine reads until newline or timeout expires
// TODO: implement timeout
func (p Process) ReadLine(timeout time.Duration) ([]byte, error) {
	b, err := ReadTill(p.Stdout, p.maxLen, '\n')
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Interactive sets the file descriptors to os.Stderr os.Stdout and os.Stdin
func (p Process) Interactive() error {
	return interactive(p, os.Stdin, os.Stdout, os.Stderr)
}

// the actual implementation of Process.Interactive
func interactive(p Process, in io.Reader, out, err io.Writer) error {
	errChan := make(chan error)
	go copyChan(p.Stdin, in, errChan)
	go copyChan(out, p.Stdout, errChan)
	go copyChan(err, p.Stderr, errChan)

	if err := p.Wait(); err != nil {
		errChan <- err
	}

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

// os/exec.Cmd methods

// StdinPipe returns a pipe that will be connected to the command's
// standard input when the command starts.
// The pipe will be closed automatically after Wait sees the command exit.
// A caller need only call Close to force the pipe to close sooner.
// For example, if the command being run will not exit until standard input
// is closed, the caller must close the pipe.
func (p *Process) StdinPipe() (io.WriteCloser, error) { return p.cmd.StdinPipe() }

// StderrPipe returns a pipe that will be connected to the command's
// standard error when the command starts.
//
// Wait will close the pipe after seeing the command exit, so most callers
// need not close the pipe themselves; however, an implication is that
// it is incorrect to call Wait before all reads from the pipe have completed.
// For the same reason, it is incorrect to use Run when using StderrPipe.
// See the StdoutPipe example for idiomatic usage.
func (p *Process) StderrPipe() (io.ReadCloser, error) { return p.cmd.StderrPipe() }

// StdoutPipe returns a pipe that will be connected to the command's
// standard output when the command starts.
//
// Wait will close the pipe after seeing the command exit, so most callers
// need not close the pipe themselves; however, an implication is that
// it is incorrect to call Wait before all reads from the pipe have completed.
// For the same reason, it is incorrect to call Run when using StdoutPipe.
// See the example for idiomatic usage.
func (p *Process) StdoutPipe() (io.ReadCloser, error) { return p.cmd.StdoutPipe() }

// Wait waits for the command to exit and waits for any copying to
// stdin or copying from stdout or stderr to complete.
//
// The command must have been started by Start.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
//
// If the command fails to run or doesn't complete successfully, the
// error is of type *ExitError. Other error types may be
// returned for I/O problems.
//
// If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits
// for the respective I/O loop copying to or from the process to complete.
//
// Wait releases any resources associated with the Cmd.
// NOTE: this is the Wait method for cmd.Wait NOT cmd.Process.Wait
func (p *Process) Wait() error { return p.cmd.Wait() }

// os.Process methods

// Kill causes the Process to exit immediately. Kill does not wait until
// the Process has actually exited. This only kills the Process itself,
// not any other processes it may have started.
func (p *Process) Kill() error { return p.cmd.Process.Kill() }

// Release releases any resources associated with the Process p,
// rendering it unusable in the future.
// Release only needs to be called if Wait is not.
func (p *Process) Release() error { return p.cmd.Process.Release() }

// Signal sends a signal to the Process.
// Sending Interrupt on Windows is not implemented.
func (p *Process) Signal(sig os.Signal) error { return p.cmd.Process.Signal(sig) }

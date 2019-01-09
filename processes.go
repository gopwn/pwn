package pwn

import (
	"io"
	"os"
	"os/exec"
	"time"
)

type ErrTimeout struct {
	// underlying error
	err string
}

func (e ErrTimeout) Error() string { return e.err }

// Spawn spawns a new process and returns it
func Spawn(path string, args ...string) (Process, error) {
	cmd := exec.Command(path, args...)

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
	data := ToBytes(t)
	// add the newline, we are "WriteLine" after all!
	data = append(data, '\n')

	// write the data to the processes standard input
	n, err := p.Stdin.Write(data)
	if err != nil {
		return err
	}
	// make sure we wrote all the data
	if n < len(data) {
		return io.ErrShortWrite
	}

	return nil
}

// ReadLine reads until newline or timeout expiers
// TODO: implement timeout
func (p Process) ReadLine(timeout time.Duration) ([]byte, error) {
	//return nil, ErrTimeout{"ReadLine: timeout exceeded"}
	b, err := ReadTill(p.Stdout, p.maxLen, '\n')
	if err != nil {
		return nil, err
	}

	return b, nil
}

// os/exec.Cmd methods
// TODO: Add documentation (aka copypaste from exec.Cmd docs)

func (p *Process) StdinPipe() (io.WriteCloser, error) { return p.cmd.StdinPipe() }
func (p *Process) StderrPipe() (io.ReadCloser, error) { return p.cmd.StderrPipe() }
func (p *Process) StdoutPipe() (io.ReadCloser, error) { return p.cmd.StdoutPipe() }

// NOTE: this is the Wait method for cmd.Wait NOT cmd.Process.Wait()
func (p *Process) Wait() error { return p.cmd.Wait() }

// os.Process methods

func (p *Process) Kill() error                { return p.cmd.Process.Kill() }
func (p *Process) Release() error             { return p.cmd.Process.Release() }
func (p *Process) Signal(sig os.Signal) error { return p.cmd.Process.Signal(sig) }

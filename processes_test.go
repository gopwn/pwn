package pwn

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestEcho(t *testing.T) {
	t.Parallel()
	expected := []byte("Hello, world!")
	p, err := Spawn("echo", "Hello, world!")
	if err != nil {
		t.Fatal(err)
	}

	output, err := p.ReadLine(time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// now make sure we got what we expected
	if !bytes.Equal(output, expected) {
		t.Fatalf("want %q got %q", expected, output)
	}
}

func TestSh(t *testing.T) {
	expected := []byte("Hello, world")
	p, err := Spawn("sh")
	if err != nil {
		t.Fatal(err)
	}

	err = p.WriteLine("echo Hello, world")
	if err != nil {
		t.Fatal(err)
	}

	out, err := p.ReadLine(time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// now check that we got the expected output
	if !bytes.Equal(out, expected) {
		t.Fatalf("want %q got %q", expected, out)
	}
}

func TestInteractive(t *testing.T) {
	// program that is used for testing.
	args := []string{"go", "run", "testdata/cat.go"}

	tests := []struct {
		stdin  string // stdin to give the process
		stderr string // expected stderr
		stdout string // expected stdout

		wantErr bool // Do we expect an error?
	}{
		{
			stdin:   "hey there",
			stderr:  "",
			stdout:  "hey there",
			wantErr: false,
		},
		{
			stdin:   "foo",
			stderr:  "",
			stdout:  "foo",
			wantErr: false,
		},
	}
	for _, tc := range tests {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		stdin := strings.NewReader(tc.stdin)

		p, err := Spawn(args[0], args[1:]...)
		if err != nil {
			t.Fatal(err)
		}

		err = interactive(p, stdin, stdout, stderr)
		if tc.wantErr && err != nil {
			t.Fatalf("wantErr: %v; got %v", tc.wantErr, err)
		}

		// now we compare the outputs to what we expect
		if stdout.String() != tc.stdout {
			t.Fatalf("want stdout: %q; got stdout %q", tc.stdout, stdout.String())
		}

		if stderr.String() != tc.stderr {
			t.Fatalf("want stderr: %q; got stderr %q", tc.stderr, stderr.String())
		}
	}
}

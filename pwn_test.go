package pwn

import "testing"

// Used for testing automated covarage, will be removed
func TestDummy(t *testing.T) {
	s := Dummy()
	if s != "dumb" {
		t.Fatalf("want: %q; got: %q", "dumb", s)
	}
}

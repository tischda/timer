package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Inspired by https://talks.golang.org/2014/testing.slide#23
func TestUsage(t *testing.T) {

	if os.Getenv("BE_CRASHER") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestUsage")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	// capture output of process execution
	r, w, _ := os.Pipe()
	cmd.Stderr = w
	err := cmd.Run()
	w.Close() //nolint:errcheck

	// check return code
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatalf("Expected exit status 1, but was: %v, ", err)
	}

	// now check that Usage message is displayed
	captured, _ := io.ReadAll(r)
	actual := string(captured)
	expected := "Usage:"

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

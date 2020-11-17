package system

import (
	"strconv"
	"testing"
)

func TestMemory(t *testing.T) {
	mem := Memory()
	if mem == "" {
		t.Fatal("expected memory size, got nothing")
	}
	// check for unit
	_, err := strconv.Atoi(mem[:len(mem)-1])
	if err != nil {
		t.Fatal(err)
	}
}

func TestPid(t *testing.T) {
	pids := Pid()
	if pids == "" {
		t.Fatal("expected pids to return number, got empty string")
	}
}

package commands

import (
	"testing"
)

func TestNormalizeCommand(t *testing.T) {
	var input string
	var expected string
	var result string

	input = "ls"
	expected = "ls"
	result = normalizeCommand(input)
	if result != expected {
		t.Fatalf("TestNormalizeCommand(\"%v\"): expected \"%v\", got \"%v\"", input, expected, result)
	}

	input = "  ls    "
	expected = "ls"
	result = normalizeCommand(input)
	if result != expected {
		t.Fatalf("TestNormalizeCommand(\"%v\"): expected \"%v\", got \"%v\"", input, expected, result)
	}

	input = "ls       -l"
	expected = "ls -l"
	result = normalizeCommand(input)
	if result != expected {
		t.Fatalf("TestNormalizeCommand(\"%v\"): expected \"%v\", got \"%v\"", input, expected, result)
	}
}

func TestRunCommand(t *testing.T) {
	var input string
	var err error

	input = "echo TestRunCommand: You should see this print"
	err = Run(input)
	if err != nil {
		t.Fatalf("TestRunCommand(%v): %v", input, err)
	}

	input = "foo"
	err = Run(input)
	if err == nil {
		t.Fatalf("TestRunCommand(%v): should have failed on bad command", input)
	}

	input = ""
	err = Run(input)
	if err == nil {
		t.Fatalf("TestRunCommand(%v): should have failed on empty command", input)
	}
}

func TestRunCommandOutput(t *testing.T) {
	var input string
	var expected string
	var output string
	var err error

	input = "echo foo"
	expected = "foo\n"
	output, err = Output(input)
	if err != nil {
		t.Fatalf("TestRunCommand(%v): %v", input, err)
	}
	if output != expected {
		t.Fatalf("TestRunCommandOutput(\"%v\"): expected \"%v\", got \"%v\"", input, expected, output)
	}
}

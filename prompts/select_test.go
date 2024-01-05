package prompts

import (
	"testing"
)

func TestSelect(t *testing.T) {
	var message string
	var options []string
	var descriptions []string
	var err error

	message = ""
	options = []string{}
	descriptions = []string{}
	_, _, err = NewSelect().SelectTitle(message).SelectOptions(options).SelectDetails(descriptions).SelectRun()
	if err == nil {
		t.Fatalf("TestSelect(\"%v\",%v,%v): should have failed because of empty options", message, options, descriptions)
	}
}

func TestMultiSelect(t *testing.T) {
	var message string
	var options []string
	var descriptions []string
	var err error

	message = ""
	options = []string{}
	descriptions = []string{}
	_, _, err = NewSelect().SelectTitle(message).SelectOptions(options).SelectDetails(descriptions).SelectMultiRun()
	if err == nil {
		t.Fatalf("TestMultiSelect(\"%v\",%v,%v): should have failed because of empty options", message, options, descriptions)
	}
}

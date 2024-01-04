package prompts

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/GrantFBarnes/go-cli/ansi"
)

type Confirm struct {
	message   string
	defaultNo bool
}

// NewConfirm returns a new Confirm object
func NewConfirm(message string) *Confirm {
	return &Confirm{
		message:   message,
		defaultNo: false,
	}
}

// MessageDefaultNo sets the default of a Confirmation prompt to no
func (c *Confirm) MessageDefaultNo() *Confirm {
	c.defaultNo = true
	return c
}

// MessageRun will prompt the user with the Confirmation prompt
func (c *Confirm) MessageRun() (bool, error) {
	fmt.Print(c.message)
	if c.defaultNo {
		fmt.Print("[y/N]")
	} else {
		fmt.Print("[Y/n]")
	}

	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadByte()
	if err != nil {
		return false, err
	}

	switch result {
	// Y
	case 121, 89:
		return true, nil
	// N
	case 110, 78:
		return false, nil
	}

	return !c.defaultNo, nil
}

type Text struct {
	message  string
	confirm  bool
	secret   bool
	required bool
}

// NewText returns a new Text object
func NewText(message string) *Text {
	return &Text{
		message:  message,
		confirm:  false,
		secret:   false,
		required: false,
	}
}

// TextConfirm sets Text prompt to confirm input
func (t *Text) TextConfirm() *Text {
	t.confirm = true
	return t
}

// TextSecret sets Text prompt to hide input
func (t *Text) TextSecret() *Text {
	t.secret = true
	return t
}

// TextRequired sets Text prompt to require input
func (t *Text) TextRequired() *Text {
	t.required = true
	return t
}

// TextRun will prompt the user with the Text prompt
func (t *Text) TextRun() (string, error) {
	fmt.Print(t.message)
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	result = result[:len(result)-1]
	if t.secret {
		ansi.CursorPreviousLine()
		ansi.EraseLine()
	}

	if t.confirm {
		fmt.Print("Again:")
		confirm, err := reader.ReadString('\n')
		confirm = confirm[:len(confirm)-1]
		if err != nil {
			return "", err
		}

		if t.secret {
			ansi.CursorPreviousLine()
			ansi.EraseLine()
		}

		if result != confirm {
			return "", errors.New("confirmation doesn't match")
		}
	}

	if t.required && len(result) == 0 {
		return "", errors.New("input is required")
	}

	return string(result), nil
}

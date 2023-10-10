package prompts

import (
	"bufio"
	"fmt"
	"os"

	"github.com/GrantFBarnes/go-cli/ansi"
)

// Prompt with message, return typed response.
func Text(message string) (string, error) {
	return prompt(message, false)
}

// Prompt with message, return typed response. Input does not display in plain text.
func Secret(message string) (string, error) {
	return prompt(message, true)
}

// Confirm yes/no. Returns boolean true==YES, false==NO. Default YES.
func Confirm(message string) (bool, error) {
	fmt.Print(message)
	fmt.Print("[Y/n]")

	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadByte()
	if err != nil {
		return false, err
	}
	isNotN := result != 110 && result != 78
	return isNotN, nil
}

func prompt(message string, hideInput bool) (string, error) {
	fmt.Print(message)

	var result []rune
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}

		if char == 10 { // enter
			break
		}

		result = append(result, char)

		if hideInput {
			ansi.CursorBack()
			fmt.Print("*")
		}
	}
	return string(result), nil
}

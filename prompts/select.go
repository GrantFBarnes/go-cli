package prompts

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/GrantFBarnes/go-cli/ansi"
)

// Offer options to select, return option selected.
func Select(title string, options []string, details []string) (string, error) {
	indexes, err := getSelectIndexes(title, options, details, false)
	if err != nil {
		return "", err
	}

	if len(indexes) != 1 {
		return "", errors.New("selection invalid")
	}

	idx := indexes[0]
	if idx < 0 {
		return "", errors.New("index invalid")
	}
	if len(options) <= idx {
		return "", errors.New("index invalid")
	}

	return options[idx], nil
}

// Offer options to select, return array of options selected.
func MultiSelect(title string, options []string, details []string) ([]string, error) {
	var result []string
	indexes, err := getSelectIndexes(title, options, details, true)
	if err != nil {
		return result, err
	}

	for _, idx := range indexes {
		if idx < 0 {
			continue
		}
		if len(options) <= idx {
			continue
		}
		result = append(result, options[idx])
	}

	return result, nil
}

func getSelectIndexes(title string, options []string, details []string, multi bool) ([]int, error) {
	var result []int

	// error if no options provided
	if len(options) == 0 {
		return result, errors.New("options are empty")
	}

	// print title if provided
	if len(title) > 0 {
		ansi.FontBold(true)
		ansi.FontUnderline(true)
		fmt.Println(title)
		ansi.FontReset()
	}

	// print empty lines for options to position cursor
	for range options {
		fmt.Println()
	}

	var currentIndex int = 0
	var selectedIndexes []bool = make([]bool, len(options))
	if !multi {
		selectedIndexes[0] = true
	}
	printOptions(options, details, currentIndex, selectedIndexes, multi)

	for {
		key, err := getKeypressMotion()
		if err != nil {
			return result, err
		}

		if key == "submit" {
			break
		}

		switch key {
		case "up":
			currentIndex -= 1
		case "down":
			currentIndex += 1
		case "select":
			selectedIndexes[currentIndex] = !selectedIndexes[currentIndex]
		case "exit":
			return result, errors.New("no selection made")
		default:
			continue
		}

		if currentIndex < 0 {
			currentIndex = len(options) - 1
		} else if currentIndex >= len(options) {
			currentIndex = 0
		}

		printOptions(options, details, currentIndex, selectedIndexes, multi)
	}

	if multi {
		for i, v := range selectedIndexes {
			if v {
				result = append(result, i)
			}
		}
	} else {
		result = append(result, currentIndex)
	}

	return result, nil
}

// Prints options for selection
func printOptions(options []string, details []string, currentIndex int, selectedIndexes []bool, multi bool) {
	ansi.CursorPreviousLines(len(options))
	for i, option := range options {
		if multi {
			if selectedIndexes[i] {
				ansi.FontColorGreen()
				fmt.Print(" [X] ")
				ansi.FontReset()
			} else {
				fmt.Print(" [ ] ")
			}
		} else {
			if i == currentIndex {
				fmt.Print(" > ")
			} else {
				fmt.Print("   ")
			}
		}

		if i == currentIndex {
			ansi.FontBold(true)
			ansi.FontColorCyan()
		}

		fmt.Print(option)

		if len(details) > i {
			if len(details[i]) > 0 {
				ansi.FontFaint(true)
				fmt.Print(" - ", details[i])
			}
		}

		ansi.FontReset()
		fmt.Println()
	}
}

// Gets single keypress from user to determine motion.
//
// returns either "submit", "up", "down", "select", "exit", or ""
func getKeypressMotion() (string, error) {
	// read single keypress
	bytes := make([]byte, 4)
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.Read(bytes)
	if err != nil {
		return "", errors.New("could not read keypress")
	}

	key, _ := utf8.DecodeRune(bytes)

	if key == 10 { // enter
		ansi.CursorPreviousLine()
		return "submit", nil
	}

	// clear line to not display keypress
	ansi.CursorLineStart()
	ansi.EraseLine()

	if key == 27 { // special
		if bytes[1] == 0 && bytes[2] == 0 { // escape
			key = 113 // set to q
		} else if bytes[1] == 91 && bytes[2] == 65 { // up arrow
			key = 107 // set to k
		} else if bytes[1] == 91 && bytes[2] == 66 { // down arrow
			key = 106 // set to j
		}
	}

	switch key {
	case 106: // j
		return "down", nil
	case 107: // k
		return "up", nil
	case 32: // space
		return "select", nil
	case 113: // q
		return "exit", nil
	default:
		return "", nil
	}
}

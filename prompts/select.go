package prompts

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/GrantFBarnes/go-cli/ansi"
	"github.com/GrantFBarnes/go-cli/commands"
)

type Select struct {
	// input parameters
	title          string
	options        []string
	details        []string
	defaultIndex   int
	maxRowsPerPage int
	eraseAfter     bool

	// calculated parameters
	rowsPerPage   int
	lastPageIndex int
}

// NewSelect returns a new Select object
func NewSelect() *Select {
	return &Select{
		title:          "",
		options:        []string{},
		details:        []string{},
		defaultIndex:   0,
		maxRowsPerPage: 20,
		eraseAfter:     false,

		rowsPerPage:   0,
		lastPageIndex: 0,
	}
}

// SelectTitle sets the title of a Selection
func (s *Select) SelectTitle(title string) *Select {
	s.title = title
	return s
}

// SelectOptions sets the options of a Selection
func (s *Select) SelectOptions(options []string) *Select {
	for _, option := range options {
		s.options = append(s.options, option)
	}
	return s.selectRowsPerPage()
}

// SelectOption adds the option to the options of a Selection
func (s *Select) SelectOption(option string) *Select {
	s.options = append(s.options, option)
	return s.selectRowsPerPage()
}

// SelectDetails sets the options of a Selection
func (s *Select) SelectDetails(details []string) *Select {
	for _, detail := range details {
		s.details = append(s.details, detail)
	}
	return s.selectRowsPerPage()
}

// SelectDetail adds the option to the details of a Selection
func (s *Select) SelectDetail(detail string) *Select {
	s.details = append(s.details, detail)
	return s.selectRowsPerPage()
}

// SelectDefaultIndex sets the default index of a Selection
func (s *Select) SelectDefaultIndex(idx int) *Select {
	s.defaultIndex = idx
	return s
}

// SelectMaxRowsPerPage sets the max rows per page of a Selection
func (s *Select) SelectMaxRowsPerPage(count int) *Select {
	if count < 1 {
		s.maxRowsPerPage = 1
	} else {
		s.maxRowsPerPage = count
	}
	return s.selectRowsPerPage()
}

// SelectEraseAfter sets the Selection to erase output after selection is made
func (s *Select) SelectEraseAfter() *Select {
	s.eraseAfter = true
	return s
}

// selectRowsPerPage will calculate the rows per page on a Selection
// based on the amount of options and max rows per page
func (s *Select) selectRowsPerPage() *Select {
	if len(s.options) < s.maxRowsPerPage {
		s.rowsPerPage = len(s.options)
	} else {
		s.rowsPerPage = s.maxRowsPerPage
	}
	return s.selectLastPageIndex()
}

// selectLastPageIndex will calculate the last page index on a Selection
// based on the amount of options and rows per page
func (s *Select) selectLastPageIndex() *Select {
	if s.rowsPerPage > 0 {
		s.lastPageIndex = (len(s.options) - 1) / s.rowsPerPage
	} else {
		s.lastPageIndex = 0
	}
	return s
}

// SelectRun to prompt user with Selection
func (s *Select) SelectRun() (int, string, error) {
	indexes, err := s.runSelection(false)
	if err != nil {
		return 0, "", err
	}

	if len(indexes) != 1 {
		return 0, "", errors.New("selection invalid")
	}

	idx := indexes[0]

	if idx < 0 {
		return 0, "", errors.New("index invalid")
	}

	if len(s.options) <= idx {
		return 0, "", errors.New("index invalid")
	}

	return idx, s.options[idx], nil
}

// SelectMultiRun to prompt user with Multi-Selection
func (s *Select) SelectMultiRun() ([]int, []string, error) {
	var resultIndexes []int
	var resultStrings []string

	indexes, err := s.runSelection(true)
	if err != nil {
		return resultIndexes, resultStrings, err
	}

	for _, idx := range indexes {
		if idx < 0 {
			continue
		}
		if len(s.options) <= idx {
			continue
		}
		resultIndexes = append(resultIndexes, idx)
		resultStrings = append(resultStrings, s.options[idx])
	}

	return resultIndexes, resultStrings, nil
}

// runSelection will run a Selection prompt, erasing after if needed
func (s *Select) runSelection(multi bool) ([]int, error) {
	result, err := s.getSelectIndexes(multi)
	if s.eraseAfter {
		linesToErase := s.rowsPerPage
		if len(s.title) > 0 {
			linesToErase += 1
		}
		if s.lastPageIndex > 0 {
			linesToErase += 1
		}
		for i := 0; i < linesToErase; i++ {
			ansi.CursorPreviousLine()
			ansi.EraseLine()
		}
	}
	return result, err
}

// getSelectIndexes gets the selected indexes of a Selection
func (s *Select) getSelectIndexes(multi bool) ([]int, error) {
	if len(s.options) == 0 {
		return nil, errors.New("no options provided")
	}

	if s.defaultIndex >= len(s.options) {
		return nil, errors.New("default index out of range")
	}

	if len(s.title) > 0 {
		ansi.FontBold(true)
		ansi.FontUnderline(true)
		fmt.Println(s.title)
		ansi.FontReset()
	}

	currentIndex := s.defaultIndex
	var selectedIndexes []bool = make([]bool, len(s.options))

	for i := 0; i < s.rowsPerPage; i++ {
		fmt.Println()
	}
	if s.lastPageIndex > 0 {
		fmt.Println()
	}

	s.printOptions(currentIndex, selectedIndexes, multi)

	for {
		key, err := getKeypressMotion()
		if err != nil {
			return nil, err
		}

		if key == "submit" {
			break
		}

		switch key {
		case "up":
			if currentIndex == 0 {
				currentIndex = len(s.options) - 1
			} else {
				currentIndex -= 1
			}
		case "down":
			currentIndex += 1
			if currentIndex >= len(s.options) {
				currentIndex = 0
			}
		case "left":
			if s.lastPageIndex == 0 {
				continue
			}

			currentPage := currentIndex / s.rowsPerPage
			if currentPage == 0 {
				currentIndex = s.lastPageIndex * s.rowsPerPage
			} else {
				currentIndex = (currentPage - 1) * s.rowsPerPage
			}
		case "right":
			if s.lastPageIndex == 0 {
				continue
			}

			currentPage := currentIndex / s.rowsPerPage
			if currentPage == s.lastPageIndex {
				currentIndex = 0
			} else {
				currentIndex = (currentPage + 1) * s.rowsPerPage
			}
		case "select":
			if multi {
				selectedIndexes[currentIndex] = !selectedIndexes[currentIndex]
			}
		case "all":
			if multi {
				allSelected := true
				for i := range selectedIndexes {
					if !selectedIndexes[i] {
						allSelected = false
						break
					}
				}
				for i := range selectedIndexes {
					selectedIndexes[i] = !allSelected
				}
			}
		case "exit":
			return nil, errors.New("no selection made")
		default:
			continue
		}

		s.printOptions(currentIndex, selectedIndexes, multi)
	}

	var result []int
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

// printOptions prints options for selection
func (s *Select) printOptions(currentIndex int, selectedIndexes []bool, multi bool) {
	for i := 0; i < s.rowsPerPage; i++ {
		ansi.CursorPreviousLine()
		ansi.EraseLine()
	}
	if s.lastPageIndex > 0 {
		ansi.CursorPreviousLine()
		ansi.EraseLine()
	}

	skip := (currentIndex / s.rowsPerPage) * s.rowsPerPage

	for i := 0; i < s.rowsPerPage; i++ {
		idx := i + skip
		if len(s.options) <= idx {
			fmt.Println()
			continue
		}

		if multi {
			if selectedIndexes[idx] {
				ansi.FontColorGreen()
				fmt.Print(" [X] ")
				ansi.FontReset()
			} else {
				fmt.Print(" [ ] ")
			}
		} else {
			if idx == currentIndex {
				fmt.Print(" > ")
			} else {
				fmt.Print("   ")
			}
		}

		if idx == currentIndex {
			ansi.FontBold(true)
			ansi.FontColorCyan()
		}

		fmt.Print(s.options[idx])

		if len(s.details) > idx {
			if len(s.details[idx]) > 0 {
				ansi.FontFaint(true)
				fmt.Print(" - ", s.details[idx])
			}
		}

		ansi.FontReset()
		fmt.Println()
	}

	if s.lastPageIndex > 0 {
		ansi.FontColorWhite()
		ansi.FontFaint(true)
		ansi.FontItalic(true)
		currentPage := currentIndex / s.rowsPerPage
		fmt.Printf("Page [%d/%d]\n", currentPage+1, s.lastPageIndex+1)
		ansi.FontReset()
	}
}

// getKeypressMotion gets single keypress from user to determine motion.
func getKeypressMotion() (string, error) {
	var err error

	err = commands.RunSilent("stty -F /dev/tty cbreak min 1")
	if err != nil {
		commands.RunSilent("stty -F /dev/tty sane")
		return "", errors.New("stty command failed")
	}

	// read single keypress
	bytes := make([]byte, 4)
	reader := bufio.NewReader(os.Stdin)
	_, err = reader.Read(bytes)
	if err != nil {
		return "", errors.New("could not read keypress")
	}

	commands.RunSilent("stty -F /dev/tty sane")

	key, _ := utf8.DecodeRune(bytes)

	// clear line to not display keypress
	ansi.CursorLineStart()
	ansi.EraseLine()

	if key == 27 { // special
		if bytes[1] == 0 && bytes[2] == 0 { // escape
			key = 113 // set to q
		} else if bytes[1] == 91 && bytes[2] == 90 { // shift tab
			key = 107 // set to k
		} else if bytes[1] == 91 && bytes[2] == 65 { // up arrow
			key = 107 // set to k
		} else if bytes[1] == 91 && bytes[2] == 66 { // down arrow
			key = 106 // set to j
		} else if bytes[1] == 91 && bytes[2] == 68 { // left arrow
			key = 104 // set to h
		} else if bytes[1] == 91 && bytes[2] == 67 { // right arrow
			key = 108 // set to l
		}
	}

	switch key {
	case 10: // enter
		ansi.CursorPreviousLine()
		return "submit", nil
	case 107: // k
		return "up", nil
	case 106: // j
		return "down", nil
	case 104: // h
		return "left", nil
	case 108: // l
		return "right", nil
	case 9: // tab
		return "down", nil
	case 32: // space
		return "select", nil
	case 97: // a
		return "all", nil
	case 113: // q
		return "exit", nil
	}
	return "", nil
}

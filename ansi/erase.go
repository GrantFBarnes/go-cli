package ansi

import "fmt"

// Erase display from cursor to end.
func EraseDisplayToEnd() {
	fmt.Print("\033[J")
}

// Erase display from cursor to beginning.
func EraseDisplayToBeginning() {
	fmt.Print("\033[1J")
}

// Erase entire display and place cursor at upper left.
func EraseDisplay() {
	fmt.Print("\033[2J")
}

// Erase line from cursor to end.
func EraseLineToEnd() {
	fmt.Print("\033[K")
}

// Erase line from cursor to beginning.
func EraseLineToBeginning() {
	fmt.Print("\033[1K")
}

// Erase enire line, cursor position does not change.
func EraseLine() {
	fmt.Print("\033[2K")
}

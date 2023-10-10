package ansi

import "fmt"

// Moves cursor up one row.
func CursorUp() {
	fmt.Print("\033[A")
}

// Moves cursor down one row.
func CursorDown() {
	fmt.Print("\033[B")
}

// Moves cursor forward one column.
func CursorForward() {
	fmt.Print("\033[C")
}

// Moves cursor back one column.
func CursorBack() {
	fmt.Print("\033[D")
}

// Moves cursor to beginning of the next line.
func CursorNextLine() {
	fmt.Print("\033[E")
}

// Moves cursor to beginning of the next n lines.
func CursorNextLines(n int) {
	fmt.Printf("\033[%vE", n)
}

// Moves cursor to beginning of the previous line.
func CursorPreviousLine() {
	fmt.Print("\033[F")
}

// Moves cursor to beginning of the previous n lines.
func CursorPreviousLines(n int) {
	fmt.Printf("\033[%vF", n)
}

// Moves the cursor to start of current line.
func CursorLineStart() {
	fmt.Print("\033[G")
}

// Moves the cursor to row n, column m. The values are 1-based.
func CursorPosition(n int, m int) {
	fmt.Printf("\033[%v;%vH", n, m)
}

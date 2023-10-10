package ansi

import "fmt"

// Reset font to normal. All attributes turned off.
func FontReset() {
	fmt.Print("\033[0m")
}

// Set font bold attribute.
func FontBold(enable bool) {
	if enable {
		fmt.Print("\033[1m")
	} else {
		fmt.Print("\033[22m")
	}
}

// Set font faint attribute.
func FontFaint(enable bool) {
	if enable {
		fmt.Print("\033[2m")
	} else {
		fmt.Print("\033[22m")
	}
}

// Set font italic attribute.
func FontItalic(enable bool) {
	if enable {
		fmt.Print("\033[3m")
	} else {
		fmt.Print("\033[23m")
	}
}

// Set font underline attribute.
func FontUnderline(enable bool) {
	if enable {
		fmt.Print("\033[4m")
	} else {
		fmt.Print("\033[24m")
	}
}

// Set font blink slowly attribute.
func FontSlowBlink(enable bool) {
	if enable {
		fmt.Print("\033[5m")
	} else {
		fmt.Print("\033[25m")
	}
}

// Set font blink rapidly attribute.
func FontRapidBlink(enable bool) {
	if enable {
		fmt.Print("\033[6m")
	} else {
		fmt.Print("\033[25m")
	}
}

// Set font invert color attribute. (swap foreground and background colors)
func FontInvertColor(enable bool) {
	if enable {
		fmt.Print("\033[7m")
	} else {
		fmt.Print("\033[27m")
	}
}

// Set font hide attribute.
func FontHide(enable bool) {
	if enable {
		fmt.Print("\033[8m")
	} else {
		fmt.Print("\033[28m")
	}
}

// Set font crossed-out/sriked attribute.
func FontStrike(enable bool) {
	if enable {
		fmt.Print("\033[9m")
	} else {
		fmt.Print("\033[29m")
	}
}

// Set font to primary (default) font.
func FontDefault() {
	fmt.Print("\033[10m")
}

// Set font color to black.
func FontColorBlack() {
	fmt.Print("\033[30m")
}

// Set font color to red.
func FontColorRed() {
	fmt.Print("\033[31m")
}

// Set font color to green.
func FontColorGreen() {
	fmt.Print("\033[32m")
}

// Set font color to yellow.
func FontColorYellow() {
	fmt.Print("\033[33m")
}

// Set font color to blue.
func FontColorBlue() {
	fmt.Print("\033[34m")
}

// Set font color to magenta.
func FontColorMagenta() {
	fmt.Print("\033[35m")
}

// Set font color to cyan.
func FontColorCyan() {
	fmt.Print("\033[36m")
}

// Set font color to white.
func FontColorWhite() {
	fmt.Print("\033[37m")
}

// Set font background color to black.
func FontBackColorBlack() {
	fmt.Print("\033[40m")
}

// Set font background color to red.
func FontBackColorRed() {
	fmt.Print("\033[41m")
}

// Set font background color to green.
func FontBackColorGreen() {
	fmt.Print("\033[42m")
}

// Set font background color to yellow.
func FontBackColorYellow() {
	fmt.Print("\033[43m")
}

// Set font background color to blue.
func FontBackColorBlue() {
	fmt.Print("\033[44m")
}

// Set font background color to magenta.
func FontBackColorMagenta() {
	fmt.Print("\033[45m")
}

// Set font background color to cyan.
func FontBackColorCyan() {
	fmt.Print("\033[46m")
}

// Set font background color to white.
func FontBackColorWhite() {
	fmt.Print("\033[47m")
}

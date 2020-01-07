package panel

import "fmt"

/*
Down below lie parts of ANSI/VT100 Terminal Control Escape Sequences, TC stands for Terminal Control.
*/
const (
	/* Cursor sequences sets the cursor position where subsequent text will begin. */
	// The cursor will move to the home position, at the upper left of the screen.
	TCCursorHome = "\033[H"
	// The cursor will move to the start of current line.
	TCCursorLineStart = "\033[;0H"

	// Erases the screen from the current line down to the bottom of the screen.
	TCEraseDown = "\033[J"
	// Erases the entire current line.
	TCEraseLine = "\033[2K"
	// Erases from the current cursor position to the start of the current line.
	TCEraseStartofLine = "\033[1K"
	// Erases from the current cursor position to the end of the current line.
	TCEraseEndofLine = "\033[K"
)

// eraseScreen clears the screen and sets the cursor at the upper left of the screen.
func eraseScreen() {
	fmt.Print(TCCursorHome)
	fmt.Print(TCEraseDown)
}

// eraseCurrentLine erases the entire current line and sets the cursor at the start of current line.
func eraseCurrentLine() {
	fmt.Print(TCEraseLine)
	fmt.Print(TCCursorLineStart)
}

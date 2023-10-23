package models

import "github.com/gdamore/tcell"

type Note struct {
	// The time in milliseconds when the note should be played.
	Time int
	// The key that should be pressed to play the note.
	Key string
	// The X location of the note on the screen.
	X int
	// The Y location of the note on the screen.
	Y int
	//Style
	Style tcell.Style
}

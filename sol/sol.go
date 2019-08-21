// Package sol provides all the Solarized colors.
package sol

import (
	"math/rand"
	"time"
)

const (
	Base03 = "\033[1;30m"
	B03    = "\033[1;30m"
	Base02 = "\033[0;30m"
	B02    = "\033[0;30m"
	Base01 = "\033[1;32m"
	B01    = "\033[1;32m"
	Base00 = "\033[1;33m"
	B00    = "\033[1;33m"
	Base0  = "\033[1;34m"
	B0     = "\033[1;34m"
	Base1  = "\033[1;36m"
	B1     = "\033[1;36m"
	Base2  = "\033[0;37m"
	B2     = "\033[0;37m"
	Base3  = "\033[1;37m"
	B3     = "\033[1;37m"

	Yellow  = "\033[0;33m"
	Y       = "\033[0;33m"
	Orange  = "\033[1;31m"
	O       = "\033[1;31m"
	Red     = "\033[0;31m"
	R       = "\033[0;31m"
	Magenta = "\033[0;35m"
	M       = "\033[0;35m"
	Violet  = "\033[1;35m"
	V       = "\033[1;35m"
	Blue    = "\033[0;34m"
	B       = "\033[0;34m"
	Cyan    = "\033[0;36m"
	C       = "\033[0;36m"
	Green   = "\033[0;32m"
	G       = "\033[0;32m"

	Reset = "\033[0m"
	X     = "\033[0m"

	ClearScreen = "\033[2J\033[H"
	ClearLine   = "\033[2K\033[G"
	CursorOff   = "\033[?25l"
	CursorOn    = "\033[?25h"
	StrikeOut   = "\033[9m"
)

// Colors is a slice of only the Solarized colors (no bases). These
// are gauranteed to always be visible. The bases are not included because when
// randomizing from the list there is a chance one or more of the bases will
// not be visible.
var Colors = [...]string{Yellow, Orange, Red, Magenta, Violet, Blue, Cyan, Green}


// Bases is a list of all the Solarized base colors (levels of grey).
var Bases = [...]string{Base02, Base01, Base00, Base0, Base1, Base2}

// Random returns a random color form SolarizedColors.
func Random() string {
	rand.Seed(time.Now().UnixNano())
	return Colors[rand.Intn(len(Colors))]
}

// MultiSol solarizes ever rune in a string randomly from the Colors.
func Multiple(s string) string {
	var m string
	for _, r := range s {
		m += Random() + string(r)
	}
	m += Reset
	return m
}

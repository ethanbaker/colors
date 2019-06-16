package colors

import (
	"math/rand"
	"strings"
	"time"
)

const (
	SolBase03 = "\033[1;30m"
	SolBase02 = "\033[0;30m"
	SolBase01 = "\033[1;32m"
	SolBase00 = "\033[1;33m"
	SolBase0  = "\033[1;34m"
	SolBase1  = "\033[1;36m"
	SolBase2  = "\033[0;37m"
	SolBase3  = "\033[1;37m"

	SolYellow  = "\033[0;33m"
	SolOrange  = "\033[1;31m"
	SolRed     = "\033[0;31m"
	SolMagenta = "\033[0;35m"
	SolViolet  = "\033[1;35m"
	SolBlue    = "\033[0;34m"
	SolCyan    = "\033[0;36m"
	SolGreen   = "\033[0;32m"

	Reset       = "\033[0m"
	ClearScreen = "\033[2J\033[H"
	ClearLine   = "\033[2K\033[G"
	CursorOff   = "\033[?25l"
	CursorOn    = "\033[?25h"
	StrikeOut   = "\033[9m"
)

var colorReplacer *strings.Replacer
var ansiReplacer *strings.Replacer

func init() {
	colorReplacer = strings.NewReplacer(
		SolBase03, "",
		SolBase02, "",
		SolBase01, "",
		SolBase00, "",
		SolBase0, "",
		SolBase1, "",
		SolBase2, "",
		SolBase3, "",
		SolYellow, "",
		SolOrange, "",
		SolRed, "",
		SolMagenta, "",
		SolViolet, "",
		SolBlue, "",
		SolCyan, "",
		SolGreen, "",
		Reset, "",
	)

	ansiReplacer = strings.NewReplacer(
		ClearScreen, "",
		ClearLine, "",
		CursorOff, "",
		CursorOn, "",
		StrikeOut, "",
	)
}

// SolarizedColors is a slice of only the Solarized colors (no bases). These
// are gauranteed to always be visible. The bases are not included because when
// randomizing from the list there is a chance one or more of the bases will
// not be visible.
var SolarizedColors = [...]string{SolYellow, SolOrange, SolRed, SolMagenta, SolViolet, SolBlue, SolCyan, SolGreen}

// SolarizedBases is a list of all the Solarized base colors (levels of grey).
var SolarizedBases = [...]string{SolBase02, SolBase01, SolBase00, SolBase0, SolBase1, SolBase2}

// RandSol returns a random color form SolarizedColors.
func RandSol() string {
	rand.Seed(time.Now().UnixNano())
	return SolarizedColors[rand.Intn(len(SolarizedColors))]
}

// MultiSol solarizes ever rune in a string randomly from the SolarizedColors.
func MultiSol(s string) string {
	var m string
	for _, r := range s {
		m += RandSol() + string(r)
	}
	m += Reset
	return m
}

// Decolor removes any of the ANSI color escapes known to this package.
func Decolor(s string) string {
	return colorReplacer.Replace(s)
}

// Strip strips any ANSI escapes known to this package.
func Strip(s string) string {
	return colorReplacer.Replace(ansiReplacer.Replace(s))
}

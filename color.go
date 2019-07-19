package color

import (
	"strings"

	"gitlab.com/ethandbaker/color/css"
	"gitlab.com/ethandbaker/color/sol"
)

var colorReplacer *strings.Replacer
var ansiReplacer *strings.Replacer

func init() {
	colorReplacer = strings.NewReplacer(
		sol.Base03, "",
		sol.Base02, "",
		sol.Base01, "",
		sol.Base00, "",
		sol.Base0, "",
		sol.Base1, "",
		sol.Base2, "",
		sol.Base3, "",
		sol.Yellow, "",
		sol.Orange, "",
		sol.Red, "",
		sol.Magenta, "",
		sol.Violet, "",
		sol.Blue, "",
		sol.Cyan, "",
		sol.Green, "",
		sol.Reset, "",
	)

	ansiReplacer = strings.NewReplacer(
		sol.ClearScreen, "",
		sol.ClearLine, "",
		sol.CursorOff, "",
		sol.CursorOn, "",
		sol.StrikeOut, "",
	)
}

// Decolor removes any of the ANSI color escapes known to this package.
func Decolor(s string) string {
	return colorReplacer.Replace(s)
}

// Strip strips any ANSI escapes known to this package.
// TODO strip all ANSI escapes
func Strip(s string) string {
	return colorReplacer.Replace(ansiReplacer.Replace(s))
}

var Index = map[string]string
  "css-red": "\033",
  "sol-red": "\033",
}

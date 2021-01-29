// Package sol provides all the Solarized colors.
package sol

import (
	"math/rand"
	"time"
)

/* Color Hex Value Constants */

const (
	// All of the solarized bases exported as decimal values
	Base01 = 0x002b36
	B01    = 0x002b36
	Base02 = 0x073642
	B02    = 0x073642
	Base03 = 0x586e75
	B03    = 0x586e75
	Base00 = 0x657b83
	B00    = 0x657b83
	Base0  = 0x839496
	B0     = 0x839496
	Base1  = 0x93a1a1
	B1     = 0x93a1a1
	Base2  = 0xeee8d5
	B2     = 0xeee8d5
	Base3  = 0xfdf6e3
	B3     = 0xfdf6e3

	// All of the solarized colors exported as decimal values
	Yellow  = 0xB58900
	Y       = 0xB58900
	Orange  = 0xCB4B16
	O       = 0xCB4B16
	Red     = 0xDC322F
	R       = 0xDC322F
	Magenta = 0xD33682
	M       = 0xD33682
	Violet  = 0x6C71C4
	V       = 0x6C71C4
	Blue    = 0x268BD2
	B       = 0x268BD2
	Cyan    = 0x2AA198
	C       = 0x2AA198
	Green   = 0x859900
	G       = 0x859900

	// Extra ansi escape codes that have no hex equivalent:
	AnsiReset = "\x1b[0m"
	AnsiX     = "\x1b[0m"
)

/* Lists of Colors */

// HexColors is a slice of only the Solarized colors (no bases) in
// hexadecimal format. These are guaranteed to always be visible. The bases
// are not included because when randomizing from the list there is a
// chance one or more of the bases will not be visible.
var HexColors = [...]int32{Yellow, Orange, Red, Magenta, Violet, Blue, Cyan, Green}

// HexBases is a slice of all the solarized base colors (levels of grey).
var HexBases = [...]int32{Base02, Base01, Base00, Base0, Base1, Base2}

// AnsiColors is an array of the solarized colors as ANSI escape codes
var AnsiColors = [...]string{AnsiIndex["yellow"], AnsiIndex["orange"], AnsiIndex["red"], AnsiIndex["magenta"], AnsiIndex["violet"], AnsiIndex["blue"], AnsiIndex["cyan"], AnsiIndex["green"]}

// AnsiBases is an array of all the solarized base colors (levels of grey).
var AnsiBases = [...]string{AnsiIndex["Base02"], AnsiIndex["Base01"], AnsiIndex["Base00"], AnsiIndex["Base0"], AnsiIndex["Base1"], AnsiIndex["Base2"]}

// ColorNames is a list of all of the solarized color names
var ColorNames = [...]string{"Yellow", "Orange", "Red", "Magenta", "Violet", "Blue", "Cyan", "Green"}

// BaseNames is a list of all of the solarized base names
var BaseNames = [...]string{"Base00", "Base01", "Base02", "Base0", "Base1", "Base2"}

/* Color Functions */

// RandomHex returns a random color as a hexadecimal value
func RandomHex() int32 {
	rand.Seed(time.Now().UnixNano())
	return HexColors[rand.Intn(len(HexColors))]
}

// RandomAnsi returns a random color as an ANSI escape code
func RandomAnsi() string {
	rand.Seed(time.Now().UnixNano())
	return AnsiIndex[ColorNames[rand.Intn(len(ColorNames))]]
}

// Multiple solarizes ever rune in a string randomly from the Colors.
func Multiple(s string) string {
	var m string
	for _, r := range s {
		m += RandomAnsi() + string(r)
	}
	m += AnsiReset
	return m
}

/* Indexes that match names to colors in different types */

// HexIndex is a map of all solarized color names and their corresponding
// decimal values
var HexIndex = map[string]int32{
	"Base01": 0x002b36,
	"B01":    0x002b36,
	"Base02": 0x073642,
	"B02":    0x073642,
	"Base03": 0x586e75,
	"B03":    0x586e75,
	"Base00": 0x657b83,
	"B00":    0x657b83,
	"Base0":  0x839496,
	"B0":     0x839496,
	"Base1":  0x93a1a1,
	"B1":     0x93a1a1,
	"Base2":  0xeee8d5,
	"B2":     0xeee8d5,
	"Base3":  0xfdf6e3,
	"B3":     0xfdf6e3,

	"Yellow":  0xB58900,
	"Y":       0xB58900,
	"Orange":  0xCB4B16,
	"O":       0xCB4B16,
	"Red":     0xDC322F,
	"R":       0xDC322F,
	"Magenta": 0xD33682,
	"M":       0xD33682,
	"Violet":  0x6C71C4,
	"V":       0x6C71C4,
	"Blue":    0x268BD2,
	"B":       0x268BD2,
	"Cyan":    0x2AA198,
	"C":       0x2AA198,
	"Green":   0x859900,
	"G":       0x859900,
}

// AnsiIndex is a map of all solarized color names and their corresponding
// ansi escape codes
var AnsiIndex = map[string]string{
	"base03": "\x1b[1;30m",
	"b03":    "\x1b[1;30m",
	"base02": "\x1b[0;30m",
	"b02":    "\x1b[0;30m",
	"base01": "\x1b[1;32m",
	"b01":    "\x1b[1;32m",
	"base00": "\x1b[1;33m",
	"b00":    "\x1b[1;33m",
	"base0":  "\x1b[1;34m",
	"b0":     "\x1b[1;34m",
	"base1":  "\x1b[1;36m",
	"b1":     "\x1b[1;36m",
	"base2":  "\x1b[0;37m",
	"b2":     "\x1b[0;37m",
	"base3":  "\x1b[1;37m",
	"b3":     "\x1b[1;37m",

	"yellow":  "\x1b[0;33m",
	"y":       "\x1b[0;33m",
	"orange":  "\x1b[1;31m",
	"o":       "\x1b[1;31m",
	"red":     "\x1b[0;31m",
	"r":       "\x1b[0;31m",
	"magenta": "\x1b[0;35m",
	"m":       "\x1b[0;35m",
	"violet":  "\x1b[1;35m",
	"v":       "\x1b[1;35m",
	"blue":    "\x1b[0;34m",
	"b":       "\x1b[0;34m",
	"cyan":    "\x1b[0;36m",
	"c":       "\x1b[0;36m",
	"green":   "\x1b[0;32m",
	"g":       "\x1b[0;32m",
}

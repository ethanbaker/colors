// Package sol provides all the Solarized colors.
package sol

import (
	"math/rand"
	"time"
)

/* Color Hex Value Constants */

const (
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

	//Ascii escape codes that have no hex equivalent:
	AsciiReset = "\033[0m"
	AsciiX     = "\033[0m"
)

/* Lists of Colors */

// HexColors is a slice of only the Solarized colors (no bases) in
// hexidecimal format. These are gauranteed to always be visible. The bases
// are not included because when randomizing from the list there is a
// chance one or more of the bases will not be visible.
var HexColors = [...]int32{Yellow, Orange, Red, Magenta, Violet, Blue, Cyan, Green}

// HexBases is a slice of all the solarized base colors (levels of grey).
var HexBases = [...]int32{Base02, Base01, Base00, Base0, Base1, Base2}

// AsciiColors is an array of the solarized colors as ascii escape codes
var AsciiColors = [...]string{AsciiIndex["yellow"], AsciiIndex["orange"], AsciiIndex["red"], AsciiIndex["magenta"], AsciiIndex["violet"], AsciiIndex["blue"], AsciiIndex["cyan"], AsciiIndex["green"]}

// AsciiBases is an array of all the solarized base colors (levels of grey).
var AsciiBases = [...]string{AsciiIndex["Base02"], AsciiIndex["Base01"], AsciiIndex["Base00"], AsciiIndex["Base0"], AsciiIndex["Base1"], AsciiIndex["Base2"]}

// ColorNames is a list of all of the solarized color names
var ColorNames = [...]string{"Yellow", "Orange", "Red", "Magenta", "Violet", "Blue", "Cyan", "Green"}

// BaseNames is a list of all of the solarized base names
var BaseNames = [...]string{"Base00", "Base01", "Base02", "Base0", "Base1", "Base2"}

/* Color Functions */

// RandomHex returns a random color as a hexidecimal value
func RandomHex() int32 {
	rand.Seed(time.Now().UnixNano())
	return HexColors[rand.Intn(len(HexColors))]
}

// RandomAscii returns a random color as an ascii escape code
func RandomAscii() string {
	rand.Seed(time.Now().UnixNano())
	return AsciiIndex[ColorNames[rand.Intn(len(ColorNames))]]
}

// Multiple solarizes ever rune in a string randomly from the Colors.
func Multiple(s string) string {
	var m string
	for _, r := range s {
		m += RandomAscii() + string(r)
	}
	m += AsciiReset
	return m
}

/* Indexes that match names to colors in different types */

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

var AsciiIndex = map[string]string{
	"base03": "\033[1;30m",
	"b03":    "\033[1;30m",
	"base02": "\033[0;30m",
	"b02":    "\033[0;30m",
	"base01": "\033[1;32m",
	"b01":    "\033[1;32m",
	"base00": "\033[1;33m",
	"b00":    "\033[1;33m",
	"base0":  "\033[1;34m",
	"b0":     "\033[1;34m",
	"base1":  "\033[1;36m",
	"b1":     "\033[1;36m",
	"base2":  "\033[0;37m",
	"b2":     "\033[0;37m",
	"base3":  "\033[1;37m",
	"b3":     "\033[1;37m",

	"yellow":  "\033[0;33m",
	"y":       "\033[0;33m",
	"orange":  "\033[1;31m",
	"o":       "\033[1;31m",
	"red":     "\033[0;31m",
	"r":       "\033[0;31m",
	"magenta": "\033[0;35m",
	"m":       "\033[0;35m",
	"violet":  "\033[1;35m",
	"v":       "\033[1;35m",
	"blue":    "\033[0;34m",
	"b":       "\033[0;34m",
	"cyan":    "\033[0;36m",
	"c":       "\033[0;36m",
	"green":   "\033[0;32m",
	"g":       "\033[0;32m",
}

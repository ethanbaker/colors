// Package sol provides all the Solarized colors.
package sol

import (
	"math/rand"
)

/* Color Hex Value Constants */

// All of the solarized bases and colors exported as ansi sequences
const (
	Base03 = "\x1b[1;30m"
	B03    = "\x1b[1;30m"
	Base02 = "\x1b[0;30m"
	B02    = "\x1b[0;30m"
	Base01 = "\x1b[1;32m"
	B01    = "\x1b[1;32m"
	Base00 = "\x1b[1;33m"
	B00    = "\x1b[1;33m"
	Base0  = "\x1b[1;34m"
	B0     = "\x1b[1;34m"
	Base1  = "\x1b[1;36m"
	B1     = "\x1b[1;36m"
	Base2  = "\x1b[0;37m"
	B2     = "\x1b[0;37m"
	Base3  = "\x1b[1;37m"
	B3     = "\x1b[1;37m"

	Base03Background = "\x1b[48;2;0;43;54m"
	B03bg            = "\x1b[48;2;0;43;54m"
	Base02Background = "\x1b[48;2;7;54;66m"
	B02bg            = "\x1b[48;2;7;54;66m"
	Base01Background = "\x1b[48;2;88;110;117m"
	B01bg            = "\x1b[48;2;88;110;117m"
	Base00Background = "\x1b[48;2;101;123;131m"
	B00bg            = "\x1b[48;2;101;123;131m"
	Base0Background  = "\x1b[48;2;131;148;150m"
	B0bg             = "\x1b[48;2;131;148;150m"
	Base1Background  = "\x1b[48;2;147;161;161m"
	B1bg             = "\x1b[48;2;147;161;161m"
	Base2Background  = "\x1b[48;2;238;232;213m"
	B2bg             = "\x1b[48;2;238;232;213m"
	Base3Background  = "\x1b[48;2;253;246;227m"
	B3bg             = "\x1b[48;2;253;246;227m"

	Yellow  = "\x1b[0;33m"
	Y       = "\x1b[0;33m"
	Orange  = "\x1b[1;31m"
	O       = "\x1b[1;31m"
	Red     = "\x1b[0;31m"
	R       = "\x1b[0;31m"
	Magenta = "\x1b[0;35m"
	M       = "\x1b[0;35m"
	Violet  = "\x1b[1;35m"
	V       = "\x1b[1;35m"
	Blue    = "\x1b[0;34m"
	B       = "\x1b[0;34m"
	Cyan    = "\x1b[0;36m"
	C       = "\x1b[0;36m"
	Green   = "\x1b[0;32m"
	G       = "\x1b[0;32m"

	YellowBackground  = "\x1b[48;2;181;137;0m"
	Ybg               = "\x1b[48;2;181;137;0m"
	OrangeBackground  = "\x1b[48;2;203;75;22m"
	Obg               = "\x1b[48;2;203;75;22m"
	RedBackground     = "\x1b[48;2;220;50;47m"
	Rbg               = "\x1b[48;2;220;50;47m"
	MagentaBackground = "\x1b[48;2;211;54;130m"
	Mbg               = "\x1b[48;2;211;54;130m"
	VioletBackground  = "\x1b[48;2;108;113;196m"
	Vbg               = "\x1b[48;2;108;113;196m"
	BlueBackground    = "\x1b[48;2;38;139;210m"
	Bbg               = "\x1b[48;2;38;139;210m"
	CyanBackground    = "\x1b[48;2;42;161;152m"
	CBg               = "\x1b[48;2;42;161;152m"
	GreenBackground   = "\x1b[48;2;133;153;0m"
	GBg               = "\x1b[48;2;133;153;0m"

	// Extra ansi escape codes
	Reset = "\x1b[0m"
	X     = "\x1b[0m"
)

// HexMap is a map of all solarized color names and their corresponding
// decimal values
var HexMap = map[string]int32{
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

// HexBaseMap is a map of all solarized base names and their corresponding
// decimal values
var HexBaseMap = map[string]int32{
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
}

// AnsiMap is a map of all solarized color names and their corresponding
// ansi escape codes
var AnsiMap = map[string]string{
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

// AnsiBackground is a map of all solarized color backgrounds and their corresponding ansi escape codes
var AnsiBackgroundMap = map[string]string{
	"yellow":  "\x1b[48;2;181;137;0m",
	"y":       "\x1b[48;2;181;137;0m",
	"orange":  "\x1b[48;2;203;75;22m",
	"o":       "\x1b[48;2;203;75;22m",
	"red":     "\x1b[48;2;220;50;47m",
	"r":       "\x1b[48;2;220;50;47m",
	"magenta": "\x1b[48;2;211;54;130m",
	"m":       "\x1b[48;2;211;54;130m",
	"violet":  "\x1b[48;2;108;113;196m",
	"v":       "\x1b[48;2;108;113;196m",
	"blue":    "\x1b[48;2;38;139;210m",
	"b":       "\x1b[48;2;38;139;210m",
	"cyan":    "\x1b[48;2;42;161;152m",
	"c":       "\x1b[48;2;42;161;152m",
	"green":   "\x1b[48;2;133;153;0m",
	"g":       "\x1b[48;2;133;153;0m",
}

// AnsiBaseMap is a map of all solarized color names and their corresponding
// ansi escape codes
var AnsiBaseMap = map[string]string{
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
}

// AnsiBaseBackgroundMap is a map of all solarized base backgrounds and their corresponding ansi escape codes
var AnsiBaseBackgroundMap = map[string]string{
	"base03": "\x1b[48;2;0;43;54m",
	"b03":    "\x1b[48;2;0;43;54m",
	"base02": "\x1b[48;2;7;54;66m",
	"b02":    "\x1b[48;2;7;54;66m",
	"base01": "\x1b[48;2;88;110;117m",
	"b01":    "\x1b[48;2;88;110;117m",
	"base00": "\x1b[48;2;101;123;131m",
	"b00":    "\x1b[48;2;101;123;131m",
	"base0":  "\x1b[48;2;131;148;150m",
	"b0":     "\x1b[48;2;131;148;150m",
	"base1":  "\x1b[48;2;147;161;161m",
	"b1":     "\x1b[48;2;147;161;161m",
	"base2":  "\x1b[48;2;238;232;213m",
	"b2":     "\x1b[48;2;238;232;213m",
	"base3":  "\x1b[48;2;253;246;227m",
	"b3":     "\x1b[48;2;253;246;227m",
}

/* Color Functions */

// Random returns a random color as an ANSI escape code
func Random() string {
	return pick(AnsiMap)
}

// RandomBackground returns a random background as an ANSI escape code
func RandomBackground() string {
	return pick(AnsiBackgroundMap)
}

// RandomBase returns a random solarized based as an ANSI escape code
func RandomBase() string {
	return pick(AnsiBaseMap)
}

// RandomBaseBackground returns a random solarized based background as an ANSI escape code
func RandomBaseBackground() string {
	return pick(AnsiBaseBackgroundMap)
}

// RandomHex returns a random color as a hexadecimal value
func RandomHex() int32 {
	return pick(HexMap)
}

// RandomHexBase returns a random base as a hexadecimal value
func RandomHexBase() int32 {
	return pick(HexBaseMap)
}

// Multiple solarizes every rune color in a string randomly from solarized colors
func Multiple(s string) string {
	var m string
	for _, r := range s {
		m += Random() + string(r)
	}
	m += Reset
	return m
}

// MultipleBase solarizes every rune color in a string randomly from solarized bases
func MultipleBase(s string) string {
	var m string
	for _, r := range s {
		m += RandomBase() + string(r)
	}
	m += Reset
	return m
}

// MultipleBackground solarizes every rune color in a string randomly from solarized background colors
func MultipleBackground(s string) string {
	var m string
	for _, r := range s {
		m += RandomBackground() + string(r)
	}
	m += Reset
	return m
}

// MultipleBaseBackground solarizes every rune color in a string randomly from solarized background base
func MultipleBaseBackground(s string) string {
	var m string
	for _, r := range s {
		m += RandomBaseBackground() + string(r)
	}
	m += Reset
	return m
}

// Helper function to choose a random element from a map
func pick[K comparable, V any](m map[K]V) V {
	k := rand.Intn(len(m))
	i := 0

	for _, x := range m {
		if i == k {
			return x
		}
		i++
	}

	panic("unreachable")
}

package color

/*
 */

import (
	"math"
	"strconv"
)

// Custom types to hold return values

// The RGB type holds three values: one for red (R), green, (G) and
// blue (B). Each of these colors are on the domain of [0, 255].
type RGB struct {
	R int `json:"R"`
	G int `json:"G"`
	B int `json:"B"`
}

// HSV type holds three values: one for the hue (H), one for the
// saturation (S), and one for the value (V). The value for hue is on the
// domain of [0, 360] and the saturation and value values are on the
// domain of [0, 100].
type HSV struct {
	H int `json:"H"`
	S int `json:"S"`
	V int `json:"V"`
}

// HSL type holds three values: one for the hue (H), one for the
// saturation (S), and one for the length (L). The value for hue is on the
// domain of [0, 360] and the saturatino and length values are on the
// domain of [0, 100].
type HSL struct {
	H int `json:"H"`
	S int `json:"S"`
	L int `json:"L"`
}

// CMYK type holds three values: one for the cyan percentage (C),
// magenta percentage (M), one for the yellow percentage (Y), and one for
// the black percentage (K). All values are on the domain [0, 100].
type CMYK struct {
	C int `json:"C"`
	M int `json:"M"`
	Y int `json:"Y"`
	K int `json:"K"`
}

// Hex type holds a hexadecimal value as a string. This can be any
// hexadecimal value from 000000 to FFFFFF. The '#' is left off for
// simplicity's sake, as it can be added later easily by the user.
type Hex string

// Decimal type holds a converted hexadecimal value to a computer-readable
// decimal value (EX: 0xFFFFFF). The converted decimal value is on the
// domain [0, 16777215].
type Decimal int

// Ansi type holds an ansi escape code. It uses the escape code format
// \x1b[38;2;R;G;Bm, where R, G, and B are the RGB values, respectively.
// Each value is on the domain [0, 255], but will always work as long as
// the input value is in the correct domain.
type Ansi string

// Color type holds all of the different color types. This is useful if
// a program is returning a certian color in all of its different types
type Color struct {
	RGB     `json:"RGB"`
	HSV     `json:"HSV"`
	HSL     `json:"HSL"`
	CMYK    `json:"CMYK"`
	Hex     `json:"Hex"`
	Decimal `json:"Decimal"`
	Ansi    `json:"Ansi"`
}

// TYPEtoRGB conversion functions

// HSVtoRGB converts HSV values to RGB values
func HSVtoRGB(hsv HSV) RGB {
	s := float64(hsv.S) / float64(100)
	v := float64(hsv.V) / float64(100)
	c := s * v
	x := c * float64(1-math.Abs(float64(math.Mod(float64(hsv.H)/float64(60), 2)-1)))
	m := v - c
	var r, g, b float64
	switch {
	case hsv.H < 60:
		r = c
		g = x
		b = 0
	case hsv.H < 120 && hsv.H >= 60:
		r = x
		g = c
		b = 0
	case hsv.H < 180 && hsv.H >= 120:
		r = 0
		g = c
		b = x
	case hsv.H < 240 && hsv.H >= 180:
		r = 0
		g = x
		b = c
	case hsv.H < 300 && hsv.H >= 240:
		r = x
		g = 0
		b = c
	case hsv.H < 360 && hsv.H >= 300:
		r = c
		g = 0
		b = x
	}
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))
	return RGB{R, G, B}
}

// HSLtoRGB converts HSL values to RGB values
func HSLtoRGB(hsl HSL) RGB {
	s := float64(hsl.S) / float64(100)
	l := float64(hsl.L) / float64(100)
	c := (1 - math.Abs(2*l-float64(1))) * s
	x := c * (float64(1) - math.Abs(math.Mod(float64(hsl.H)/60, 2)-float64(1)))
	m := l - c/2
	var r, g, b float64
	switch {
	case hsl.H < 60:
		r = c
		g = x
		b = 0
	case hsl.H < 120 && hsl.H >= 60:
		r = x
		g = c
		b = 0
	case hsl.H < 180 && hsl.H >= 120:
		r = 0
		g = c
		b = x
	case hsl.H < 240 && hsl.H >= 180:
		r = 0
		g = x
		b = c
	case hsl.H < 300 && hsl.H >= 240:
		r = x
		g = 0
		b = c
	case hsl.H < 360 && hsl.H >= 300:
		r = c
		g = 0
		b = x
	}
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))
	return RGB{R, G, B}
}

// HextoRGB converts a hexadecimal string to RGB values
func HextoRGB(hex Hex) RGB {
	if hex[0:1] == "#" {
		hex = hex[1:]
	}
	r := string(hex)[0:2]
	g := string(hex)[2:4]
	b := string(hex)[4:6]
	R, _ := strconv.ParseInt(r, 16, 0)
	G, _ := strconv.ParseInt(g, 16, 0)
	B, _ := strconv.ParseInt(b, 16, 0)

	return RGB{int(R), int(G), int(B)}
}

// DecimaltoRGB converts a decimal value to RGB values
func DecimaltoRGB(decimal Decimal) RGB {
	hex := Ansi(strconv.FormatInt(int64(decimal), 16))
	for len(hex) < 6 {
		hex = "0" + hex
	}
	rgb := HextoRGB(Hex(hex))

	return RGB{int(rgb.R), int(rgb.G), int(rgb.B)}
}

// CMYKtoRGB converts CMYK values to RGB Values
func CMYKtoRGB(cmyk CMYK) RGB {
	c := float64(cmyk.C) / 100
	m := float64(cmyk.M) / 100
	y := float64(cmyk.Y) / 100
	k := float64(cmyk.K) / 100
	R := int(math.Round(255 * (1 - c) * (1 - k)))
	G := int(math.Round(255 * (1 - m) * (1 - k)))
	B := int(math.Round(255 * (1 - y) * (1 - k)))
	return RGB{R, G, B}
}

//RGBtoTYPE conversion functions

// RGBtoHSV converts RGB values to HSV values
func RGBtoHSV(rgb RGB) HSV {
	r := float64(rgb.R) / 255
	g := float64(rgb.G) / 255
	b := float64(rgb.B) / 255

	var cmax float64
	switch {
	case r >= g && r >= b:
		cmax = r
	case g >= r && g >= b:
		cmax = g
	case b >= r && b >= g:
		cmax = b
	}

	var cmin float64
	switch {
	case r <= g && r <= b:
		cmin = r
	case g <= r && g <= b:
		cmin = g
	case b <= r && b <= g:
		cmin = b
	}

	d := cmax - cmin

	var H float64
	switch {
	case cmax == 0 || d == 0:
		H = 0
	case r == cmax:
		H = math.Round(60 * math.Mod((g-b)/d, 6))
	case g == cmax:
		H = math.Round(60 * ((b-r)/d + 2))
	case b == cmax:
		H = math.Round(60 * ((r-g)/d + 4))
	}
	if H < 0 {
		H += 360
	}

	var S float64
	if cmax == 0 {
		S = 0
	} else {
		S = math.Round(d / cmax * 100)
	}

	V := math.Round(cmax * 100)

	return HSV{int(H), int(S), int(V)}
}

// RGBtoHSL converts RGB values to HSL values
func RGBtoHSL(rgb RGB) HSL {
	return HSVtoHSL(RGBtoHSV(rgb))
}

// RGBtoHex converts RGB values to a hexadecimal value in a string
func RGBtoHex(rgb RGB) Hex {
	hex := strconv.FormatInt(int64(rgb.R*65536+rgb.G*256+rgb.B), 16)
	for len(hex) < 6 {
		hex = "0" + hex
	}

	return Hex(hex)
}

// RGBtoDecimal converts RGB values to a decimal value
func RGBtoDecimal(rgb RGB) Decimal {
	decimal := rgb.R*65536 + rgb.G*256 + rgb.B

	return Decimal(decimal)
}

// RGBtoCMYK converts RGB values to CMYK values
func RGBtoCMYK(rgb RGB) CMYK {
	r := float64(rgb.R) / 255
	g := float64(rgb.G) / 255
	b := float64(rgb.B) / 255

	var k float64
	switch {
	case r >= g && r >= b:
		k = 1 - r
	case g >= r && g >= b:
		k = 1 - g
	case b >= r && b >= g:
		k = 1 - b
	}

	var c, m, y float64
	if k != 1 {
		c = (1 - r - k) / (1 - k)
		m = (1 - g - k) / (1 - k)
		y = (1 - b - k) / (1 - k)
	} else {
		c = 0
		m = 0
		y = 0
	}
	return CMYK{int(math.Round(c * 100)), int(math.Round(m * 100)), int(math.Round(y * 100)), int(math.Round(k * 100))}
}

// Other type conversions using RGB as a standard (except HSVtoHSL for
// simplicity and convience)

// HSVtoHSL converts HSV values to HSL values
func HSVtoHSL(hsv HSV) HSL {
	s := float64(hsv.S) / float64(100)
	v := float64(hsv.V) / float64(100)
	l := float64(2-s) * float64(v/2)

	if l != 0 {
		switch {
		case l == 1:
			s = 0
		case l < 0.5:
			s = s * v / (l * 2)
		default:
			s = s * v / (2 - l*2)
		}
	}

	return HSL{hsv.H, int(math.Round(s * 100)), int(math.Round(l * 100))}
}

// HSVtoCMYK converts HSV values to CMYK values
func HSVtoCMYK(hsv HSV) CMYK {
	return RGBtoCMYK(HSVtoRGB(hsv))
}

// HSVtoHex converts HSV values to a hexadecimal string
func HSVtoHex(hsv HSV) Hex {
	return RGBtoHex(HSVtoRGB(hsv))
}

// HSVtoDecimal converts HSV values to a decimal value
func HSVtoDecimal(hsv HSV) Decimal {
	return RGBtoDecimal(HSVtoRGB(hsv))
}

// HSLtoHSV converts HSL values to HSV values
func HSLtoHSV(hsl HSL) HSV {
	return RGBtoHSV(HSLtoRGB(hsl))
}

// HSLtoCMYK converts HSL values to CMYK values
func HSLtoCMYK(hsl HSL) CMYK {
	return RGBtoCMYK(HSLtoRGB(hsl))
}

// HSLtoHex converts HSL values to a hexadecimal string
func HSLtoHex(hsl HSL) Hex {
	return RGBtoHex(HSLtoRGB(hsl))
}

// HSLtoDecimal converts HSL values to a decimal value
func HSLtoDecimal(hsl HSL) Decimal {
	return RGBtoDecimal(HSLtoRGB(hsl))
}

// HextoHSV converts hexadecimal string to HSV values
func HextoHSV(hex Hex) HSV {
	return RGBtoHSV(HextoRGB(hex))
}

// HextoHSL converts hexadecimal string to HSL values
func HextoHSL(hex Hex) HSL {
	return RGBtoHSL(HextoRGB(hex))
}

// HextoCMYK converts hexadecimal string to CMYK values
func HextoCMYK(hex Hex) CMYK {
	return RGBtoCMYK(HextoRGB(hex))
}

// HextoDecimal converts hexadecimal string to decimal value
func HextoDecimal(hex Hex) Decimal {
	return RGBtoDecimal(HextoRGB(hex))
}

// DecimaltoHSV converts decimal value to HSV values
func DecimaltoHSV(decimal Decimal) HSV {
	return RGBtoHSV(DecimaltoRGB(decimal))
}

// DecimaltoHSL converts decimal value to HSL values
func DecimaltoHSL(decimal Decimal) HSL {
	return RGBtoHSL(DecimaltoRGB(decimal))
}

// DecimaltoCMYK converts decimal value to CMYK values
func DecimaltoCMYK(decimal Decimal) CMYK {
	return RGBtoCMYK(DecimaltoRGB(decimal))
}

// DecimaltoHex converts decimal value to hexadecimal string
func DecimaltoHex(decimal Decimal) Hex {
	return RGBtoHex(DecimaltoRGB(decimal))
}

// CMYKtoHSV converts CMYK values to HSV values
func CMYKtoHSV(cmyk CMYK) HSV {
	return RGBtoHSV(CMYKtoRGB(cmyk))
}

// CMYKtoHSL converts CMYK values to HSL values
func CMYKtoHSL(cmyk CMYK) HSL {
	return RGBtoHSL(CMYKtoRGB(cmyk))
}

// CMYKtoHex converts CMYK values to a hexadecimal string
func CMYKtoHex(cmyk CMYK) Hex {
	return RGBtoHex(CMYKtoRGB(cmyk))
}

// CMYKtoDecimal converts CMYK values to a decimal value
func CMYKtoDecimal(cmyk CMYK) Decimal {
	return RGBtoDecimal(CMYKtoRGB(cmyk))
}

// Ansi Conversions

// RGBtoAnsi converts RGB values to an Ansi escape code
func RGBtoAnsi(rgb RGB) Ansi {
	str := "\x1b[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
	return Ansi(str)
}

// HSVtoAnsi converts HSV values to an Ansi escape code
func HSVtoAnsi(hsv HSV) Ansi {
	rgb := HSVtoRGB(hsv)
	str := "\x1b[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
	return Ansi(str)
}

// HSLtoAnsi converts HSL values to an Ansi escape code
func HSLtoAnsi(hsl HSL) Ansi {
	rgb := HSLtoRGB(hsl)
	str := "\x1b[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
	return Ansi(str)
}

// DecimaltoAnsi converts a decimal value to an Ansi escape code
func DecimaltoAnsi(decimal Decimal) Ansi {
	rgb := DecimaltoRGB(decimal)
	str := "\x1b[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
	return Ansi(str)
}

// HextoAnsi converts a hexadecimal string to an Ansi escape code
func HextoAnsi(hex Hex) Ansi {
	rgb := HextoRGB(hex)
	str := "\x1b[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
	return Ansi(str)
}

// CMYKtoAnsi converts CMYK values to an Ansi escape code
func CMYKtoAnsi(cmyk CMYK) Ansi {
	rgb := CMYKtoRGB(cmyk)
	str := "\x1b[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
	return Ansi(str)
}

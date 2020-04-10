package color

/*
* Color Format Transfersion
*
* All of the functions below convert different color types from one type
* to another. The color types that are included below are (named
* accordingly):
*			- RGB
*			- HSV
*			- HSL
*			- CYMK
*			- Hex
*
* All functions are in the format: TYPE1toTYPE2.
* Examples: HSLtoRGB(), RGBtoHex(), CMYKtoRGB()
*
* All color types can be converted to and from the standard RGB format
* (not necessarily from one to another). However, you can transform HSV
* format into HSL format and vice versa, because of the similarities
* between the two types.
*
* If you have a color in CMYK format and want to convert it to HSL, you
* would first convert the color in CMYK format into RGB format and then
* into HSL (written below in code).
*
* c, m, y, k := cpick.RGBtoCMYK(cpick.HSLtoRGB(0, 100, 100))
 */

import (
	"math"
	"strconv"
)

//TYPEtoRGB

func HSVtoRGB(h int, S int, V int) (int, int, int) {
	s := float64(S) / float64(100)
	v := float64(V) / float64(100)
	c := s * v
	x := c * float64(1-math.Abs(float64(math.Mod(float64(h)/float64(60), 2)-1)))
	m := v - c
	var r, g, b float64
	switch {
	case h < 60:
		r = c
		g = x
		b = 0
	case h < 120 && h >= 60:
		r = x
		g = c
		b = 0
	case h < 180 && h >= 120:
		r = 0
		g = c
		b = x
	case h < 240 && h >= 180:
		r = 0
		g = x
		b = c
	case h < 300 && h >= 240:
		r = x
		g = 0
		b = c
	case h < 360 && h >= 300:
		r = c
		g = 0
		b = x
	}
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))
	return R, G, B
}

func HSLtoRGB(H int, S int, L int) (int, int, int) {
	s := float64(S) / float64(100)
	l := float64(L) / float64(100)
	c := (1 - math.Abs(2*l-float64(1))) * s
	x := c * (float64(1) - math.Abs(math.Mod(float64(H)/60, 2)-float64(1)))
	m := l - c/2
	var r, g, b float64
	switch {
	case H < 60:
		r = c
		g = x
		b = 0
	case H < 120 && H >= 60:
		r = x
		g = c
		b = 0
	case H < 180 && H >= 120:
		r = 0
		g = c
		b = x
	case H < 240 && H >= 180:
		r = 0
		g = x
		b = c
	case H < 300 && H >= 240:
		r = x
		g = 0
		b = c
	case H < 360 && H >= 300:
		r = c
		g = 0
		b = x
	}
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))
	return R, G, B

}

func HexNametoRGB(hex string) (int, int, int) {
	if hex[0:1] == "#" {
		hex = hex[1:]
	}
	r := hex[0:2]
	g := hex[2:4]
	b := hex[4:6]
	R, _ := strconv.ParseInt(r, 16, 0)
	G, _ := strconv.ParseInt(g, 16, 0)
	B, _ := strconv.ParseInt(b, 16, 0)

	return int(R), int(G), int(B)
}

func HexValtoRGB(hexInt int) (int, int, int) {
	hex := strconv.FormatInt(int64(hexInt), 16)
	for len(hex) < 6 {
		hex = "0" + hex
	}
	r, g, b := HexNametoRGB(hex)

	return int(r), int(g), int(b)
}

func CMYKtoRGB(C int, M int, Y int, K int) (int, int, int) {
	c := float64(C) / 100
	m := float64(M) / 100
	y := float64(Y) / 100
	k := float64(K) / 100
	R := int(math.Round(255 * (1 - c) * (1 - k)))
	G := int(math.Round(255 * (1 - m) * (1 - k)))
	B := int(math.Round(255 * (1 - y) * (1 - k)))
	return R, G, B
}

//RGBtoTYPE
func RGBtoHSV(R int, G int, B int) (int, int, int) {
	r := float64(R) / 255
	g := float64(G) / 255
	b := float64(B) / 255

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

	var H int
	switch {
	case r == cmax:
		H = 60 * int(math.Mod((g-b)/d, 6))
	case g == cmax:
		H = 60 * int((b-r)/d+2)
	case b == cmax:
		H = 60 * int((r-g)/d+4)
	case 0 == cmax:
		H = 0
	}

	var S int
	if cmax == 0 {
		S = int(0)
	} else {
		S = int(d / cmax * 100)
	}

	V := int(cmax * 100)

	return H, S, V
}

func RGBtoHSL(R int, G int, B int) (int, int, int) {
	return HSVtoHSL(RGBtoHSV(R, G, B))
}

func RGBtoHexName(r int, g int, b int) string {
	hex := strconv.FormatInt(int64(r*65536+g*256+b), 16)
	for len(hex) < 6 {
		hex = "0" + hex
	}
	return hex
}

func RGBtoHexVal(r int, g int, b int) int {
	hex := r*65536 + g*256 + b
	return hex
}

func RGBtoCMYK(R int, G int, B int) (int, int, int, int) {
	r := float64(R) / 255
	g := float64(G) / 255
	b := float64(B) / 255

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
	return int(math.Round(c * 100)), int(math.Round(m * 100)), int(math.Round(y * 100)), int(math.Round(k * 100))
}

//HSVtoHSL and HSLtoHSV (for similarity of HSV and HSL and convinience sake)
func HSVtoHSL(H int, S int, V int) (int, int, int) {
	s := float64(S) / float64(100)
	v := float64(V) / float64(100)
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

	return H, int(math.Round(s * 100)), int(math.Round(l * 100))
}

func HSLtoHSV(H int, S int, L int) (int, int, int) {
	return RGBtoHSV(HSLtoRGB(H, S, L))
}

//Other type conversions using RGB as a standard
func HSVtoCMYK(h int, s int, v int) (int, int, int, int) {
	return RGBtoCMYK(HSVtoRGB(h, s, v))
}

func HSVtoHexName(h int, s int, v int) string {
	return RGBtoHexName(HSVtoRGB(h, s, v))
}

func HSVtoHexVal(h int, s int, v int) int {
	return RGBtoHexVal(HSVtoRGB(h, s, v))
}

func HSLtoCMYK(h int, s int, l int) (int, int, int, int) {
	return RGBtoCMYK(HSLtoRGB(h, s, l))
}

func HSLtoHexName(h int, s int, l int) string {
	return RGBtoHexName(HSLtoRGB(h, s, l))
}

func HSLtoHexVal(h int, s int, l int) int {
	return RGBtoHexVal(HSLtoRGB(h, s, l))
}

func HexNametoHSV(hex string) (int, int, int) {
	return RGBtoHSV(HexNametoRGB(hex))
}

func HexNametoHSL(hex string) (int, int, int) {
	return RGBtoHSL(HexNametoRGB(hex))
}

func HexNametoCMYK(hex string) (int, int, int, int) {
	return RGBtoCMYK(HexNametoRGB(hex))
}

func HexNametoHexVal(hex string) int {
	return RGBtoHexVal(HexNametoRGB(hex))
}

func HexValtoHSV(hex int) (int, int, int) {
	return RGBtoHSV(HexValtoRGB(hex))
}

func HexValtoHSL(hex int) (int, int, int) {
	return RGBtoHSL(HexValtoRGB(hex))
}

func HexValtoCMYK(hex int) (int, int, int, int) {
	return RGBtoCMYK(HexValtoRGB(hex))
}

func HexValtoHexName(hex int) string {
	return RGBtoHexName(HexValtoRGB(hex))
}

func CMYKtoHSV(c int, m int, y int, k int) (int, int, int) {
	return RGBtoHSV(CMYKtoRGB(c, m, y, k))
}

func CMYKtoHSL(c int, m int, y int, k int) (int, int, int) {
	return RGBtoHSL(CMYKtoRGB(c, m, y, k))
}

func CMYKtoHexName(c int, m int, y int, k int) string {
	return RGBtoHexName(CMYKtoRGB(c, m, y, k))
}

func CMYKtoHexVal(c int, m int, y int, k int) int {
	return RGBtoHexVal(CMYKtoRGB(c, m, y, k))
}

//Ascii Conversions
func RGBtoAscii(r int, g int, b int) string {
	str := "\033[38;2;" + string(r) + ";" + string(g) + ";" + string(b) + "m"
	return str
}

func HSVtoAscii(h int, s int, v int) string {
	r, g, b := HSVtoRGB(h, s, v)
	str := "\033[38;2;" + string(r) + ";" + string(g) + ";" + string(b) + "m"
	return str
}

func HSLtoAscii(h int, s int, l int) string {
	r, g, b := HSLtoRGB(h, s, l)
	str := "\033[38;2;" + string(r) + ";" + string(g) + ";" + string(b) + "m"
	return str
}

func HexValtoAscii(hex int) string {
	r, g, b := HexValtoRGB(hex)
	str := "\033[38;2;" + strconv.FormatInt(int64(r), 10) + ";" + strconv.FormatInt(int64(g), 10) + ";" + strconv.FormatInt(int64(b), 10) + "m"
	return str
}

func HexNametoAscii(hex string) string {
	r, g, b := HexNametoRGB(hex)
	str := "\033[38;2;" + string(r) + ";" + string(g) + ";" + string(b) + "m"
	return str
}

func CMYKtoAscii(c int, m int, y int, k int) string {
	r, g, b := CMYKtoRGB(c, m, y, k)
	str := "\033[38;2;" + string(r) + ";" + string(g) + ";" + string(b) + "m"
	return str
}

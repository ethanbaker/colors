package color_test

import (
	"testing"

	color "github.com/ethanbaker/colors"
	"github.com/stretchr/testify/assert"
)

// Test TYPEtoRGB
func TestHSVtoRGB(t *testing.T) {
	assert := assert.New(t)

	var rgb color.RGB

	rgb = color.HSVtoRGB(color.HSV{H: 0, S: 0, V: 0}) // black
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for black")
	assert.Equal(0, rgb.G, "RGB.G value was not translated correctly for black")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for black")

	rgb = color.HSVtoRGB(color.HSV{H: 0, S: 0, V: 100}) // white
	assert.Equal(255, rgb.R, "RGB.R value was not translated correctly for white")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for white")
	assert.Equal(255, rgb.B, "RGB.B value was not translated correctly for white")

	rgb = color.HSVtoRGB(color.HSV{H: 70, S: 50, V: 100}) // lime
	assert.Equal(234, rgb.R, "RGB.R value was not translated correctly for lime")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for lime")
	assert.Equal(128, rgb.B, "RGB.B value was not translated correctly for lime")

	rgb = color.HSVtoRGB(color.HSV{H: 120, S: 100, V: 50}) // green
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly green")
	assert.Equal(128, rgb.G, "RGB.G value was not translated correctly green")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly green")

	rgb = color.HSVtoRGB(color.HSV{H: 190, S: 20, V: 40}) // dark cyan
	assert.Equal(82, rgb.R, "RGB.R value was not translated correctly for dark cyan")
	assert.Equal(99, rgb.G, "RGB.G value was not translated correctly for dark cyan")
	assert.Equal(102, rgb.B, "RGB.B value was not translated correctly for dark cyan")

	rgb = color.HSVtoRGB(color.HSV{H: 260, S: 90, V: 50}) // purple
	assert.Equal(51, rgb.R, "RGB.R value was not translated correctly for purple")
	assert.Equal(13, rgb.G, "RGB.G value was not translated correctly for purple")
	assert.Equal(128, rgb.B, "RGB.B value was not translated correctly for purple")

	rgb = color.HSVtoRGB(color.HSV{H: 350, S: 75, V: 30}) // maroon
	assert.Equal(77, rgb.R, "RGB.R value was not translated correctly for maroon")
	assert.Equal(19, rgb.G, "RGB.G value was not translated correctly for maroon")
	assert.Equal(29, rgb.B, "RGB.B value was not translated correctly for maroon")
}

func TestHSLtoRGB(t *testing.T) {
	assert := assert.New(t)

	var rgb color.RGB

	rgb = color.HSLtoRGB(color.HSL{H: 0, S: 0, L: 0}) // black
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for black")
	assert.Equal(0, rgb.G, "RGB.G value was not translated correctly for black")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for black")

	rgb = color.HSLtoRGB(color.HSL{H: 0, S: 0, L: 100}) // white
	assert.Equal(255, rgb.R, "RGB.R value was not translated correctly for white")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for white")
	assert.Equal(255, rgb.B, "RGB.B value was not translated correctly for white")

	rgb = color.HSLtoRGB(color.HSL{H: 70, S: 100, L: 75}) // lime
	assert.Equal(234, rgb.R, "RGB.R value was not translated correctly for lime")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for lime")
	assert.Equal(128, rgb.B, "RGB.B value was not translated correctly for lime")

	rgb = color.HSLtoRGB(color.HSL{H: 120, S: 100, L: 25}) // green
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for green")
	assert.Equal(128, rgb.G, "RGB.G value was not translated correctly for green")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for green")

	rgb = color.HSLtoRGB(color.HSL{H: 190, S: 11, L: 36}) // dark cyan
	assert.Equal(82, rgb.R, "RGB.R value was not translated correctly for dark cyan")
	assert.Equal(99, rgb.G, "RGB.G value was not translated correctly for dark cyan")
	assert.Equal(102, rgb.B, "RGB.B value was not translated correctly for dark cyan")

	rgb = color.HSLtoRGB(color.HSL{H: 260, S: 82, L: 22}) // purple
	assert.Equal(41, rgb.R, "RGB.R value was not translated correctly for purple")
	assert.Equal(10, rgb.G, "RGB.G value was not translated correctly for purple")
	assert.Equal(102, rgb.B, "RGB.B value was not translated correctly for purple")

	rgb = color.HSLtoRGB(color.HSL{H: 350, S: 60, L: 19}) // maroon
	assert.Equal(78, rgb.R, "RGB.R value was not translated correctly for maroon")
	assert.Equal(19, rgb.G, "RGB.G value was not translated correctly for maroon")
	assert.Equal(29, rgb.B, "RGB.B value was not translated correctly for maroon")
}

func TestHextoRGB(t *testing.T) {
	assert := assert.New(t)

	var rgb color.RGB

	rgb = color.HextoRGB(color.Hex("#000000")) // black
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for black")
	assert.Equal(0, rgb.G, "RGB.G value was not translated correctly for black")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for black")

	rgb = color.HextoRGB(color.Hex("#ffffff")) // white
	assert.Equal(255, rgb.R, "RGB.R value was not translated correctly for white")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for white")
	assert.Equal(255, rgb.B, "RGB.B value was not translated correctly for white")

	rgb = color.HextoRGB(color.Hex("eaff80")) // lime
	assert.Equal(234, rgb.R, "RGB.R value was not translated correctly for lime")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for lime")
	assert.Equal(128, rgb.B, "RGB.B value was not translated correctly for lime")

	rgb = color.HextoRGB(color.Hex("008000")) // green
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for green")
	assert.Equal(128, rgb.G, "RGB.G value was not translated correctly for green")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for green")

	rgb = color.HextoRGB(color.Hex("#526366")) // dark cyan
	assert.Equal(82, rgb.R, "RGB.R value was not translated correctly for dark cyan")
	assert.Equal(99, rgb.G, "RGB.G value was not translated correctly for dark cyan")
	assert.Equal(102, rgb.B, "RGB.B value was not translated correctly for dark cyan")

	rgb = color.HextoRGB(color.Hex("#290A66")) // purple
	assert.Equal(41, rgb.R, "RGB.R value was not translated correctly for purple")
	assert.Equal(10, rgb.G, "RGB.G value was not translated correctly for purple")
	assert.Equal(102, rgb.B, "RGB.B value was not translated correctly for purple")

	rgb = color.HextoRGB(color.Hex("4E131D")) // maroon
	assert.Equal(78, rgb.R, "RGB.R value was not translated correctly for maroon")
	assert.Equal(19, rgb.G, "RGB.G value was not translated correctly for maroon")
	assert.Equal(29, rgb.B, "RGB.B value was not translated correctly for maroon")
}

func TestDecimaltoRGB(t *testing.T) {
	assert := assert.New(t)

	var rgb color.RGB

	rgb = color.DecimaltoRGB(color.Decimal(0)) // black
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for black")
	assert.Equal(0, rgb.G, "RGB.G value was not translated correctly for black")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for black")

	rgb = color.DecimaltoRGB(color.Decimal(16777215)) // white
	assert.Equal(255, rgb.R, "RGB.R value was not translated correctly for white")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for white")
	assert.Equal(255, rgb.B, "RGB.B value was not translated correctly for white")

	rgb = color.DecimaltoRGB(color.Decimal(32768)) // green
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for green")
	assert.Equal(128, rgb.G, "RGB.G value was not translated correctly for green")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for green")
}

func TestCMYKtoRGB(t *testing.T) {
	assert := assert.New(t)

	var rgb color.RGB

	rgb = color.CMYKtoRGB(color.CMYK{C: 0, M: 0, Y: 0, K: 100}) // black
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for black")
	assert.Equal(0, rgb.G, "RGB.G value was not translated correctly for black")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for black")

	rgb = color.CMYKtoRGB(color.CMYK{C: 0, M: 0, Y: 0, K: 0}) // white
	assert.Equal(255, rgb.R, "RGB.R value was not translated correctly for white")
	assert.Equal(255, rgb.G, "RGB.G value was not translated correctly for white")
	assert.Equal(255, rgb.B, "RGB.B value was not translated correctly for white")

	rgb = color.CMYKtoRGB(color.CMYK{C: 100, M: 0, Y: 100, K: 50}) // green
	assert.Equal(0, rgb.R, "RGB.R value was not translated correctly for green")
	assert.Equal(128, rgb.G, "RGB.G value was not translated correctly for green")
	assert.Equal(0, rgb.B, "RGB.B value was not translated correctly for green")
}

// Test RGBtoTYPE
func TestRGBtoHSV(t *testing.T) {
	assert := assert.New(t)

	var hsv color.HSV

	hsv = color.RGBtoHSV(color.RGB{R: 0, G: 0, B: 0}) // black
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for black")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for black")
	assert.Equal(0, hsv.V, "HSV.V value was not translated correctly for black")

	hsv = color.RGBtoHSV(color.RGB{R: 255, G: 255, B: 255}) // white
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for white")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for white")
	assert.Equal(100, hsv.V, "HSV.V value was not translated correctly for white")

	hsv = color.RGBtoHSV(color.RGB{R: 234, G: 255, B: 128}) // lime
	assert.Equal(70, hsv.H, "HSV.H value was not translated correctly for lime")
	assert.Equal(50, hsv.S, "HSV.S value was not translated correctly for lime")
	assert.Equal(100, hsv.V, "HSV.V value was not translated correctly for lime")

	hsv = color.RGBtoHSV(color.RGB{R: 0, G: 128, B: 0}) // green
	assert.Equal(120, hsv.H, "HSV.H value was not translated correctly for green")
	assert.Equal(100, hsv.S, "HSV.S value was not translated correctly for green")
	assert.Equal(50, hsv.V, "HSV.V value was not translated correctly for green")

	hsv = color.RGBtoHSV(color.RGB{R: 82, G: 99, B: 102}) // dark cyan
	assert.Equal(189, hsv.H, "HSV.H value was not translated correctly for dark cyan")
	assert.Equal(20, hsv.S, "HSV.S value was not translated correctly for dark cyan")
	assert.Equal(40, hsv.V, "HSV.V value was not translated correctly for dark cyan")

	hsv = color.RGBtoHSV(color.RGB{R: 51, G: 13, B: 128}) // purple
	assert.Equal(260, hsv.H, "HSV.H value was not translated correctly for purple")
	assert.Equal(90, hsv.S, "HSV.S value was not translated correctly for purple")
	assert.Equal(50, hsv.V, "HSV.V value was not translated correctly for purple")

	hsv = color.RGBtoHSV(color.RGB{R: 77, G: 19, B: 29}) // maroon
	assert.Equal(350, hsv.H, "HSV.H value was not translated correctly for maroon")
	assert.Equal(75, hsv.S, "HSV.S value was not translated correctly for maroon")
	assert.Equal(30, hsv.V, "HSV.V value was not translated correctly for maroon")
}

func TestRGBtoHSL(t *testing.T) {
	assert := assert.New(t)

	var hsl color.HSL

	hsl = color.RGBtoHSL(color.RGB{R: 0, G: 0, B: 0}) // black
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for black")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for black")
	assert.Equal(0, hsl.L, "HSL.L value was not translated correctly for black")

	hsl = color.RGBtoHSL(color.RGB{R: 255, G: 255, B: 255}) // white
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for white")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for white")
	assert.Equal(100, hsl.L, "HSL.L value was not translated correctly for white")

	hsl = color.RGBtoHSL(color.RGB{R: 0, G: 128, B: 0}) // green
	assert.Equal(120, hsl.H, "HSL.H value was not translated correctly for green")
	assert.Equal(100, hsl.S, "HSL.S value was not translated correctly for green")
	assert.Equal(25, hsl.L, "HSL.L value was not translated correctly for green")
}

func TestRGBtoHex(t *testing.T) {
	assert := assert.New(t)

	var hex color.Hex

	hex = color.RGBtoHex(color.RGB{R: 0, G: 0, B: 0}) // black
	assert.Equal(color.Hex("000000"), hex, "Hex value was not translated correctly for black")

	hex = color.RGBtoHex(color.RGB{R: 255, G: 255, B: 255}) // white
	assert.Equal(color.Hex("ffffff"), hex, "Hex value was not translated correctly for white")

	hex = color.RGBtoHex(color.RGB{R: 0, G: 128, B: 0}) // green
	assert.Equal(color.Hex("008000"), hex, "Hex value was not translated correctly for white")
}

func TestRGBtoDecimal(t *testing.T) {
	assert := assert.New(t)

	var decimal color.Decimal

	decimal = color.RGBtoDecimal(color.RGB{R: 0, G: 0, B: 0}) // black
	assert.Equal(color.Decimal(0), decimal, "Decimal value was not translated correctly for black")

	decimal = color.RGBtoDecimal(color.RGB{R: 255, G: 255, B: 255}) // white
	assert.Equal(color.Decimal(16777215), decimal, "Decimal value was not translated correctly for white")

	decimal = color.RGBtoDecimal(color.RGB{R: 0, G: 128, B: 0}) // green
	assert.Equal(color.Decimal(32768), decimal, "Decimal value was not translated correctly for green")
}

func TestRGBtoCMYK(t *testing.T) {
	assert := assert.New(t)

	var cmyk color.CMYK

	cmyk = color.RGBtoCMYK(color.RGB{R: 0, G: 0, B: 0}) // black
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for black")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for black")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for black")
	assert.Equal(100, cmyk.K, "CMYK.K value was not translated correctly for black")

	cmyk = color.RGBtoCMYK(color.RGB{R: 255, G: 255, B: 255}) // white
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for white")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for white")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for white")
	assert.Equal(0, cmyk.K, "CMYK.K value was not translated correctly for white")

	cmyk = color.RGBtoCMYK(color.RGB{R: 0, G: 128, B: 0}) // green
	assert.Equal(100, cmyk.C, "CMYK.C value was not translated correctly for green")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for green")
	assert.Equal(100, cmyk.Y, "CMYK.Y value was not translated correctly for green")
	assert.Equal(50, cmyk.K, "CMYK.K value was not translated correctly for green")

	cmyk = color.RGBtoCMYK(color.RGB{R: 51, G: 13, B: 128}) // purple
	assert.Equal(60, cmyk.C, "CMYK.C value was not translated correctly for purple")
	assert.Equal(90, cmyk.M, "CMYK.M value was not translated correctly for purple")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for purple")
	assert.Equal(50, cmyk.K, "CMYK.K value was not translated correctly for purple")
}

// Test HSVtoHSL and HSLtoHSV
func TestHSVtoHSL(t *testing.T) {
	assert := assert.New(t)

	var hsl color.HSL

	hsl = color.HSVtoHSL(color.HSV{H: 0, S: 0, V: 0}) // black
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for black")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for black")
	assert.Equal(0, hsl.L, "HSL.L value was not translated correctly for black")

	hsl = color.HSVtoHSL(color.HSV{H: 0, S: 0, V: 100}) // white
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for white")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for white")
	assert.Equal(100, hsl.L, "HSL.L value was not translated correctly for white")

	hsl = color.HSVtoHSL(color.HSV{H: 70, S: 50, V: 100}) // lime
	assert.Equal(70, hsl.H, "HSL.H value was not translated correctly for white")
	assert.Equal(100, hsl.S, "HSL.S value was not translated correctly for white")
	assert.Equal(75, hsl.L, "HSL.L value was not translated correctly for white")

	hsl = color.HSVtoHSL(color.HSV{H: 120, S: 100, V: 50}) // green
	assert.Equal(120, hsl.H, "HSL.H value was not translated correctly for green")
	assert.Equal(100, hsl.S, "HSL.S value was not translated correctly for green")
	assert.Equal(25, hsl.L, "HSL.L value was not translated correctly for green")
}

func TestHSVtoCMYK(t *testing.T) {
	assert := assert.New(t)

	var cmyk color.CMYK

	cmyk = color.HSVtoCMYK(color.HSV{H: 0, S: 0, V: 0}) // black
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for black")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for black")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for black")
	assert.Equal(100, cmyk.K, "CMYK.K value was not translated correctly for black")

	cmyk = color.HSVtoCMYK(color.HSV{H: 0, S: 0, V: 100}) // white
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for white")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for white")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for white")
	assert.Equal(0, cmyk.K, "CMYK.K value was not translated correctly for white")

	cmyk = color.HSVtoCMYK(color.HSV{H: 120, S: 100, V: 50}) // green
	assert.Equal(100, cmyk.C, "CMYK.C value was not translated correctly for green")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for green")
	assert.Equal(100, cmyk.Y, "CMYK.Y value was not translated correctly for green")
	assert.Equal(50, cmyk.K, "CMYK.K value was not translated correctly for green")
}

func TestHSVtoHex(t *testing.T) {
	assert := assert.New(t)

	var hex color.Hex

	hex = color.HSVtoHex(color.HSV{H: 0, S: 0, V: 0}) // black
	assert.Equal(color.Hex("000000"), hex, "Hex value was not translated correctly for black")

	hex = color.HSVtoHex(color.HSV{H: 0, S: 0, V: 100}) // white
	assert.Equal(color.Hex("ffffff"), hex, "Hex value was not translated correctly for white")

	hex = color.HSVtoHex(color.HSV{H: 120, S: 100, V: 50}) // green
	assert.Equal(color.Hex("008000"), hex, "Hex value was not translated correctly for green")
}

func TestHSVtoDecimal(t *testing.T) {
	assert := assert.New(t)

	var decimal color.Decimal

	decimal = color.HSVtoDecimal(color.HSV{H: 0, S: 0, V: 0}) // black
	assert.Equal(color.Decimal(0), decimal, "Hex value was not translated correctly for black")

	decimal = color.HSVtoDecimal(color.HSV{H: 0, S: 0, V: 100}) // white
	assert.Equal(color.Decimal(16777215), decimal, "Decimal value was not translated correctly for white")

	decimal = color.HSVtoDecimal(color.HSV{H: 120, S: 100, V: 50}) // green
	assert.Equal(color.Decimal(32768), decimal, "Decimal value was not translated correctly for green")
}

func TestHSLtoHSV(t *testing.T) {
	assert := assert.New(t)

	var hsv color.HSV

	hsv = color.HSLtoHSV(color.HSL{H: 0, S: 0, L: 0}) // black
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for black")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for black")
	assert.Equal(0, hsv.V, "HSV.V value was not translated correctly for black")

	hsv = color.HSLtoHSV(color.HSL{H: 0, S: 0, L: 100}) // white
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for white")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for white")
	assert.Equal(100, hsv.V, "HSV.V value was not translated correctly for white")

	hsv = color.HSLtoHSV(color.HSL{H: 120, S: 100, L: 25}) // green
	assert.Equal(120, hsv.H, "HSV.H value was not translated correctly for green")
	assert.Equal(100, hsv.S, "HSV.S value was not translated correctly for green")
	assert.Equal(50, hsv.V, "HSV.V value was not translated correctly for green")
}

func TestHSLtoCMYK(t *testing.T) {
	assert := assert.New(t)

	var cmyk color.CMYK

	cmyk = color.HSLtoCMYK(color.HSL{H: 0, S: 0, L: 0}) // black
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for black")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for black")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for black")
	assert.Equal(100, cmyk.K, "CMYK.K value was not translated correctly for black")

	cmyk = color.HSLtoCMYK(color.HSL{H: 0, S: 0, L: 100}) // white
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for white")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for white")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for white")
	assert.Equal(0, cmyk.K, "CMYK.K value was not translated correctly for white")

	cmyk = color.HSLtoCMYK(color.HSL{H: 120, S: 100, L: 25}) // green
	assert.Equal(100, cmyk.C, "CMYK.C value was not translated correctly for green")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for green")
	assert.Equal(100, cmyk.Y, "CMYK.Y value was not translated correctly for green")
	assert.Equal(50, cmyk.K, "CMYK.K value was not translated correctly for green")
}

func TestHSLtoHex(t *testing.T) {
	assert := assert.New(t)

	var hex color.Hex

	hex = color.HSLtoHex(color.HSL{H: 0, S: 0, L: 0}) // black
	assert.Equal(color.Hex("000000"), hex, "Hex value was not translated correctly for black")

	hex = color.HSLtoHex(color.HSL{H: 0, S: 0, L: 100}) // white
	assert.Equal(color.Hex("ffffff"), hex, "Hex value was not translated correctly for white")

	hex = color.HSLtoHex(color.HSL{H: 120, S: 100, L: 25}) // green
	assert.Equal(color.Hex("008000"), hex, "Hex value was not translated correctly for white")
}

func TestHSLtoDecimal(t *testing.T) {
	assert := assert.New(t)

	var decimal color.Decimal

	decimal = color.HSLtoDecimal(color.HSL{H: 0, S: 0, L: 0}) // black
	assert.Equal(color.Decimal(0), decimal, "Decimal value was not translated correctly for black")

	decimal = color.HSLtoDecimal(color.HSL{H: 0, S: 0, L: 100}) // white
	assert.Equal(color.Decimal(16777215), decimal, "Decimal value was not translated correctly for white")

	decimal = color.HSLtoDecimal(color.HSL{H: 120, S: 100, L: 25}) // green
	assert.Equal(color.Decimal(32768), decimal, "Decimal value was not translated correctly for green")
}

func TestHextoHSV(t *testing.T) {
	assert := assert.New(t)

	var hsv color.HSV

	hsv = color.HextoHSV(color.Hex("#000000")) // black
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for black")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for black")
	assert.Equal(0, hsv.V, "HSV.V value was not translated correctly for black")

	hsv = color.HextoHSV(color.Hex("#ffffff")) // white
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for white")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for white")
	assert.Equal(100, hsv.V, "HSV.V value was not translated correctly for white")

	hsv = color.HextoHSV(color.Hex("008000")) // green
	assert.Equal(120, hsv.H, "HSV.H value was not translated correctly for green")
	assert.Equal(100, hsv.S, "HSV.S value was not translated correctly for green")
	assert.Equal(50, hsv.V, "HSV.V value was not translated correctly for green")
}

func TestHextoHSL(t *testing.T) {
	assert := assert.New(t)

	var hsl color.HSL

	hsl = color.HextoHSL(color.Hex("#000000")) // black
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for black")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for black")
	assert.Equal(0, hsl.L, "HSL.L value was not translated correctly for black")

	hsl = color.HextoHSL(color.Hex("#ffffff")) // white
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for white")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for white")
	assert.Equal(100, hsl.L, "HSL.L value was not translated correctly for white")

	hsl = color.HextoHSL(color.Hex("008000")) // green
	assert.Equal(120, hsl.H, "HSL.H value was not translated correctly for green")
	assert.Equal(100, hsl.S, "HSL.S value was not translated correctly for green")
	assert.Equal(25, hsl.L, "HSL.L value was not translated correctly for green")
}

func TestHextoCMYK(t *testing.T) {
	assert := assert.New(t)

	var cmyk color.CMYK

	cmyk = color.HextoCMYK(color.Hex("#000000")) // black
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for black")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for black")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for black")
	assert.Equal(100, cmyk.K, "CMYK.K value was not translated correctly for black")

	cmyk = color.HextoCMYK(color.Hex("#ffffff")) // white
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for white")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for white")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for white")
	assert.Equal(0, cmyk.K, "CMYK.K value was not translated correctly for white")

	cmyk = color.HextoCMYK(color.Hex("008000")) // green
	assert.Equal(100, cmyk.C, "CMYK.C value was not translated correctly for green")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for green")
	assert.Equal(100, cmyk.Y, "CMYK.Y value was not translated correctly for green")
	assert.Equal(50, cmyk.K, "CMYK.K value was not translated correctly for green")
}

func TestHextoDecimal(t *testing.T) {
	assert := assert.New(t)

	var decimal color.Decimal

	decimal = color.HextoDecimal(color.Hex("#000000")) // black
	assert.Equal(color.Decimal(0), decimal, "Decimal value was not translated correctly for black")

	decimal = color.HextoDecimal(color.Hex("#ffffff")) // white
	assert.Equal(color.Decimal(16777215), decimal, "Decimal value was not translated correctly for white")

	decimal = color.HextoDecimal(color.Hex("008000")) // green
	assert.Equal(color.Decimal(32768), decimal, "Decimal value was not translated correctly for green")
}

func TestDecimaltoHSV(t *testing.T) {
	assert := assert.New(t)

	var hsv color.HSV

	hsv = color.DecimaltoHSV(color.Decimal(0)) // black
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for black")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for black")
	assert.Equal(0, hsv.V, "HSV.V value was not translated correctly for black")

	hsv = color.DecimaltoHSV(color.Decimal(16777215)) // white
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for white")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for white")
	assert.Equal(100, hsv.V, "HSV.V value was not translated correctly for white")

	hsv = color.DecimaltoHSV(color.Decimal(32768)) // green
	assert.Equal(120, hsv.H, "HSV.H value was not translated correctly for green")
	assert.Equal(100, hsv.S, "HSV.S value was not translated correctly for green")
	assert.Equal(50, hsv.V, "HSV.V value was not translated correctly for green")
}

func TestDecimaltoHSL(t *testing.T) {
	assert := assert.New(t)

	var hsl color.HSL

	hsl = color.DecimaltoHSL(color.Decimal(0)) // black
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for black")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for black")
	assert.Equal(0, hsl.L, "HSL.L value was not translated correctly for black")

	hsl = color.DecimaltoHSL(color.Decimal(16777215)) // white
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for white")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for white")
	assert.Equal(100, hsl.L, "HSL.L value was not translated correctly for white")

	hsl = color.DecimaltoHSL(color.Decimal(32768)) // green
	assert.Equal(120, hsl.H, "HSL.H value was not translated correctly for green")
	assert.Equal(100, hsl.S, "HSL.S value was not translated correctly for green")
	assert.Equal(25, hsl.L, "HSL.L value was not translated correctly for green")
}

func TestDecimaltoCMYK(t *testing.T) {
	assert := assert.New(t)

	var cmyk color.CMYK

	cmyk = color.DecimaltoCMYK(color.Decimal(0)) // black
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for black")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for black")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for black")
	assert.Equal(100, cmyk.K, "CMYK.K value was not translated correctly for black")

	cmyk = color.DecimaltoCMYK(color.Decimal(16777215)) // white
	assert.Equal(0, cmyk.C, "CMYK.C value was not translated correctly for white")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for white")
	assert.Equal(0, cmyk.Y, "CMYK.Y value was not translated correctly for white")
	assert.Equal(0, cmyk.K, "CMYK.K value was not translated correctly for white")

	cmyk = color.DecimaltoCMYK(color.Decimal(32768)) // green
	assert.Equal(100, cmyk.C, "CMYK.C value was not translated correctly for green")
	assert.Equal(0, cmyk.M, "CMYK.M value was not translated correctly for green")
	assert.Equal(100, cmyk.Y, "CMYK.Y value was not translated correctly for green")
	assert.Equal(50, cmyk.K, "CMYK.K value was not translated correctly for green")
}

func TestDecimaltoHex(t *testing.T) {
	assert := assert.New(t)

	var hex color.Hex

	hex = color.DecimaltoHex(color.Decimal(0)) // black
	assert.Equal(color.Hex("000000"), hex, "Hex value was not translated correctly for black")

	hex = color.DecimaltoHex(color.Decimal(16777215)) // white
	assert.Equal(color.Hex("ffffff"), hex, "Hex value was not translated correctly for white")

	hex = color.DecimaltoHex(color.Decimal(32768)) // green
	assert.Equal(color.Hex("008000"), hex, "Hex value was not translated correctly for white")
}

func TestCMYKtoHSV(t *testing.T) {
	assert := assert.New(t)

	var hsv color.HSV

	hsv = color.CMYKtoHSV(color.CMYK{C: 0, M: 0, Y: 0, K: 100}) // black
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for black")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for black")
	assert.Equal(0, hsv.V, "HSV.V value was not translated correctly for black")

	hsv = color.CMYKtoHSV(color.CMYK{C: 0, M: 0, Y: 0, K: 0}) // white
	assert.Equal(0, hsv.H, "HSV.H value was not translated correctly for white")
	assert.Equal(0, hsv.S, "HSV.S value was not translated correctly for white")
	assert.Equal(100, hsv.V, "HSV.V value was not translated correctly for white")

	hsv = color.CMYKtoHSV(color.CMYK{C: 100, M: 0, Y: 100, K: 50}) // green
	assert.Equal(120, hsv.H, "HSV.H value was not translated correctly for green")
	assert.Equal(100, hsv.S, "HSV.S value was not translated correctly for green")
	assert.Equal(50, hsv.V, "HSV.V value was not translated correctly for green")
}

func TestCMYKtoHSL(t *testing.T) {
	assert := assert.New(t)

	var hsl color.HSL

	hsl = color.CMYKtoHSL(color.CMYK{C: 0, M: 0, Y: 0, K: 100}) // black
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for black")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for black")
	assert.Equal(0, hsl.L, "HSL.L value was not translated correctly for black")

	hsl = color.CMYKtoHSL(color.CMYK{C: 0, M: 0, Y: 0, K: 0}) // white
	assert.Equal(0, hsl.H, "HSL.H value was not translated correctly for white")
	assert.Equal(0, hsl.S, "HSL.S value was not translated correctly for white")
	assert.Equal(100, hsl.L, "HSL.L value was not translated correctly for white")

	hsl = color.CMYKtoHSL(color.CMYK{C: 100, M: 0, Y: 100, K: 50}) // green
	assert.Equal(120, hsl.H, "HSL.H value was not translated correctly for green")
	assert.Equal(100, hsl.S, "HSL.S value was not translated correctly for green")
	assert.Equal(25, hsl.L, "HSL.L value was not translated correctly for green")
}

func TestCMYKtoHex(t *testing.T) {
	assert := assert.New(t)

	var hex color.Hex

	hex = color.CMYKtoHex(color.CMYK{C: 0, M: 0, Y: 0, K: 100}) // black
	assert.Equal(color.Hex("000000"), hex, "Hex value was not translated correctly for black")

	hex = color.CMYKtoHex(color.CMYK{C: 0, M: 0, Y: 0, K: 0}) // white
	assert.Equal(color.Hex("ffffff"), hex, "Hex value was not translated correctly for white")

	hex = color.CMYKtoHex(color.CMYK{C: 100, M: 0, Y: 100, K: 50}) // green
	assert.Equal(color.Hex("008000"), hex, "Hex value was not translated correctly for white")
}

func TestCMYKtoDecimal(t *testing.T) {
	assert := assert.New(t)

	var decimal color.Decimal

	decimal = color.CMYKtoDecimal(color.CMYK{C: 0, M: 0, Y: 0, K: 100}) // black
	assert.Equal(color.Decimal(0), decimal, "Decimal value was not translated correctly for black")

	decimal = color.CMYKtoDecimal(color.CMYK{C: 0, M: 0, Y: 0, K: 0}) // white
	assert.Equal(color.Decimal(16777215), decimal, "Decimal value was not translated correctly for white")

	decimal = color.CMYKtoDecimal(color.CMYK{C: 100, M: 0, Y: 100, K: 50}) // green
	assert.Equal(color.Decimal(32768), decimal, "Decimal value was not translated correctly for green")
}

func TestRGBtoAnsi(t *testing.T) {
	assert := assert.New(t)

	var ansi color.Ansi

	ansi = color.RGBtoAnsi(color.RGB{R: 0, G: 0, B: 0})
	assert.Equal(color.Ansi("\x1b[38;2;0;0;0m"), ansi, "Ansi value was not translated correctly for black")

	ansi = color.RGBtoAnsi(color.RGB{R: 255, G: 255, B: 255})
	assert.Equal(color.Ansi("\x1b[38;2;255;255;255m"), ansi, "Ansi value was not translated correctly for white")

	ansi = color.RGBtoAnsi(color.RGB{R: 0, G: 128, B: 0})
	assert.Equal(color.Ansi("\x1b[38;2;0;128;0m"), ansi, "Ansi value was not translated correctly for green")
}

func TestHSVtoAnsi(t *testing.T) {
	assert := assert.New(t)

	var ansi color.Ansi

	ansi = color.HSVtoAnsi(color.HSV{H: 0, S: 0, V: 0})
	assert.Equal(color.Ansi("\x1b[38;2;0;0;0m"), ansi, "Ansi value was not translated correctly for black")

	ansi = color.HSVtoAnsi(color.HSV{H: 0, S: 0, V: 100})
	assert.Equal(color.Ansi("\x1b[38;2;255;255;255m"), ansi, "Ansi value was not translated correctly for white")

	ansi = color.HSVtoAnsi(color.HSV{H: 120, S: 100, V: 50})
	assert.Equal(color.Ansi("\x1b[38;2;0;128;0m"), ansi, "Ansi value was not translated correctly for green")
}

func TestHSLtoAnsi(t *testing.T) {
	assert := assert.New(t)

	var ansi color.Ansi

	ansi = color.HSLtoAnsi(color.HSL{H: 0, S: 0, L: 0})
	assert.Equal(color.Ansi("\x1b[38;2;0;0;0m"), ansi, "Ansi value was not translated correctly for black")

	ansi = color.HSLtoAnsi(color.HSL{H: 0, S: 0, L: 100})
	assert.Equal(color.Ansi("\x1b[38;2;255;255;255m"), ansi, "Ansi value was not translated correctly for white")

	ansi = color.HSLtoAnsi(color.HSL{H: 120, S: 100, L: 25})
	assert.Equal(color.Ansi("\x1b[38;2;0;128;0m"), ansi, "Ansi value was not translated correctly for green")
}

func TestDecimaltoAnsi(t *testing.T) {
	assert := assert.New(t)

	var ansi color.Ansi

	ansi = color.DecimaltoAnsi(color.Decimal(0))
	assert.Equal(color.Ansi("\x1b[38;2;0;0;0m"), ansi, "Ansi value was not translated correctly for black")

	ansi = color.DecimaltoAnsi(color.Decimal(16777215))
	assert.Equal(color.Ansi("\x1b[38;2;255;255;255m"), ansi, "Ansi value was not translated correctly for white")

	ansi = color.DecimaltoAnsi(color.Decimal(32768))
	assert.Equal(color.Ansi("\x1b[38;2;0;128;0m"), ansi, "Ansi value was not translated correctly for green")
}

func TestHextoAnsi(t *testing.T) {
	assert := assert.New(t)

	var ansi color.Ansi

	ansi = color.HextoAnsi(color.Hex("000000"))
	assert.Equal(color.Ansi("\x1b[38;2;0;0;0m"), ansi, "Ansi value was not translated correctly for black")

	ansi = color.HextoAnsi(color.Hex("ffffff"))
	assert.Equal(color.Ansi("\x1b[38;2;255;255;255m"), ansi, "Ansi value was not translated correctly for white")

	ansi = color.HextoAnsi(color.Hex("008000"))
	assert.Equal(color.Ansi("\x1b[38;2;0;128;0m"), ansi, "Ansi value was not translated correctly for green")
}

func TestCMYKtoAnsi(t *testing.T) {
	assert := assert.New(t)

	var ansi color.Ansi

	ansi = color.CMYKtoAnsi(color.CMYK{C: 0, M: 0, Y: 0, K: 100})
	assert.Equal(color.Ansi("\x1b[38;2;0;0;0m"), ansi, "Ansi value was not translated correctly for black")

	ansi = color.CMYKtoAnsi(color.CMYK{C: 0, M: 0, Y: 0, K: 0})
	assert.Equal(color.Ansi("\x1b[38;2;255;255;255m"), ansi, "Ansi value was not translated correctly for white")

	ansi = color.CMYKtoAnsi(color.CMYK{C: 100, M: 0, Y: 100, K: 50})
	assert.Equal(color.Ansi("\x1b[38;2;0;128;0m"), ansi, "Ansi value was not translated correctly for green")
}

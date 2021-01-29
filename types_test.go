package color_test

import (
	"fmt"
	"testing"

	"gitlab.com/ethanbakerdev/colors"
)

//Test TYPEtoRGB
func TestHSVtoRGB(t *testing.T) {
	var rgb color.RGB

	rgb = color.HSVtoRGB(color.HSV{0, 0, 0}) //black
	if rgb.R == 0 && rgb.G == 0 && rgb.B == 0 {
		fmt.Println("# Black (HSV: 0, 0, 0) is working with HSVtoRGB()")
	} else {
		t.Errorf("Black (HSV: 0, 0, 0) not working with HSVtoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.HSVtoRGB(color.HSV{0, 0, 100}) //white
	if rgb.R == 255 && rgb.G == 255 && rgb.B == 255 {
		fmt.Println("# White (HSV: 0, 0, 100) is working with HSVtoRGB()")
	} else {
		t.Errorf("White (HSV: 0, 0, 100) not working with HSVtoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.HSVtoRGB(color.HSV{120, 100, 50}) //green
	if rgb.R == 0 && rgb.G == 128 && rgb.B == 0 {
		fmt.Println("# Green (HSV: 120, 100, 50) is working with HSVtoRGB()")
	} else {
		t.Errorf("Green (HSV: 120, 100, 50) not working with HSVtoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", rgb.R, rgb.G, rgb.B)
	}
}

func TestHSLtoRGB(t *testing.T) {
	var rgb color.RGB

	rgb = color.HSLtoRGB(color.HSL{0, 0, 0}) //black
	if rgb.R == 0 && rgb.G == 0 && rgb.B == 0 {
		fmt.Println("# Black (HSL: 0, 0, 0) is working with HSLtoRGB()")
	} else {
		t.Errorf("Black (HSL: 0, 0, 0) has a problem with HSLtoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.HSLtoRGB(color.HSL{0, 0, 100}) //white
	if rgb.R == 255 && rgb.G == 255 && rgb.B == 255 {
		fmt.Println("# White (HSL: 0, 0, 100) is working with HSLtoRGB()")
	} else {
		t.Errorf("White (HSL: 0, 0, 100) has a problem with HSLtoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.HSLtoRGB(color.HSL{120, 100, 25}) //green
	if rgb.R == 0 && rgb.G == 128 && rgb.B == 0 {
		fmt.Println("# Green (HSL: 120, 100, 25) is working with HSLtoRGB()")
	} else {
		t.Errorf("Green (HSL: 120, 100, 25) has a problem with HSLtoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", rgb.R, rgb.G, rgb.B)
	}
}

func TestHextoRGB(t *testing.T) {
	var rgb color.RGB

	rgb = color.HextoRGB(color.Hex("#000000")) //black
	if rgb.R == 0 && rgb.G == 0 && rgb.B == 0 {
		fmt.Println("# Black (Hex: #000000) is working with HextoRGB()")
	} else {
		t.Errorf("Black (Hex: #000000) has a problem with HextoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.HextoRGB(color.Hex("#ffffff")) //white
	if rgb.R == 255 && rgb.G == 255 && rgb.B == 255 {
		fmt.Println("# White (Hex: #ffffff) is working with HextoRGB()")
	} else {
		t.Errorf("White (Hex: #ffffff) has a problem with HextoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.HextoRGB(color.Hex("008000")) //green
	if rgb.R == 0 && rgb.G == 128 && rgb.B == 0 {
		fmt.Println("# Green (Hex: #008000) is working with HextoRGB()")
	} else {
		t.Errorf("Green (Hex: #008000) has a problem with HextoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", rgb.R, rgb.G, rgb.B)
	}
}

func TestDecimaltoRGB(t *testing.T) {
	var rgb color.RGB

	rgb = color.DecimaltoRGB(color.Decimal(0)) //black
	if rgb.R == 0 && rgb.G == 0 && rgb.B == 0 {
		fmt.Println("# Black (Hex: #000000) is working with DecimaltoRGB()")
	} else {
		t.Errorf("Black (Hex: #000000) has a problem with DecimaltoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.DecimaltoRGB(color.Decimal(16777215)) //white
	if rgb.R == 255 && rgb.G == 255 && rgb.B == 255 {
		fmt.Println("# White (Hex: 16777215) is working with DecimaltoRGB()")
	} else {
		t.Errorf("White (Hex: 16777215) has a problem with DecimaltoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.DecimaltoRGB(color.Decimal(32768)) //green
	if rgb.R == 0 && rgb.G == 128 && rgb.B == 0 {
		fmt.Println("# Green (Hex: 32768) is working with DecimaltoRGB()")
	} else {
		t.Errorf("Green (Hex: 32768) has a problem with DecimaltoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", rgb.R, rgb.G, rgb.B)
	}
}

func TestCMYKtoRGB(t *testing.T) {
	var rgb color.RGB

	rgb = color.CMYKtoRGB(color.CMYK{0, 0, 0, 100}) //black
	if rgb.R == 0 && rgb.G == 0 && rgb.B == 0 {
		fmt.Println("# Black (CMYK: 0, 0, 0, 100) is working with CMYKtoRGB()")
	} else {
		t.Errorf("Black (CMYK: 0, 0, 0, 100) has a problem with CMYKtoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.CMYKtoRGB(color.CMYK{0, 0, 0, 0}) //white
	if rgb.R == 255 && rgb.G == 255 && rgb.B == 255 {
		fmt.Println("# White (CMYK: 0, 0, 0, 0) is working with CMYKtoRGB()")
	} else {
		t.Errorf("White (CMYK: 0, 0, 0, 100) has a problem with CMYKtoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", rgb.R, rgb.G, rgb.B)
	}

	rgb = color.CMYKtoRGB(color.CMYK{100, 0, 100, 50}) //green
	if rgb.R == 0 && rgb.G == 128 && rgb.B == 0 {
		fmt.Println("# Green (CMYK: 100, 0, 100, 50) is working with CMYKtoRGB()")
	} else {
		t.Errorf("Green (CMYK: 100, 0, 100, 50) has a problem with CMYKtoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", rgb.R, rgb.G, rgb.B)
	}
}

//Test RGBtoTYPE
func TestRGBtoHSV(t *testing.T) {
	var hsv color.HSV

	hsv = color.RGBtoHSV(color.RGB{0, 0, 0}) //black
	if hsv.H == 0 && hsv.S == 0 && hsv.V == 0 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoHSV()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoHSV()! Returned %v, %v, %v instead of 0, 0, 0.", hsv.H, hsv.S, hsv.V)
	}

	hsv = color.RGBtoHSV(color.RGB{255, 255, 255}) //white
	if hsv.H == 0 && hsv.S == 0 && hsv.V == 100 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHSV()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHSV()! Returned %v, %v, %v instead of 0, 0, 100.", hsv.H, hsv.S, hsv.V)
	}

	hsv = color.RGBtoHSV(color.RGB{0, 128, 0}) //green
	if hsv.H == 120 && hsv.S == 100 && hsv.V == 50 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHSV()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHSV()! Returned %v, %v, %v instead of 120, 100, 50.", hsv.H, hsv.S, hsv.V)
	}
}

func TestRGBtoHSL(t *testing.T) {
	var hsl color.HSL

	hsl = color.RGBtoHSL(color.RGB{0, 0, 0}) //black
	if hsl.H == 0 && hsl.S == 0 && hsl.L == 0 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoHSL()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoHSL()! Returned %v, %v, %v instead of 0, 0, 0.", hsl.H, hsl.S, hsl.L)
	}

	hsl = color.RGBtoHSL(color.RGB{255, 255, 255}) //white
	if hsl.H == 0 && hsl.S == 0 && hsl.L == 100 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHSL()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHSL()! Returned %v, %v, %v instead of 0, 0, 100.", hsl.H, hsl.S, hsl.L)
	}

	hsl = color.RGBtoHSL(color.RGB{0, 128, 0}) //green
	if hsl.H == 120 && hsl.S == 100 && hsl.L == 25 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHSL()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHSL()! Returned %v, %v, %v instead of 120, 100, 25.", hsl.H, hsl.S, hsl.L)
	}
}

func TestRGBtoHex(t *testing.T) {
	var hex color.Hex

	hex = color.RGBtoHex(color.RGB{0, 0, 0}) //black
	if hex == "000000" {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoHex()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoHex()! Returned %v instead of 000000.", hex)
	}

	hex = color.RGBtoHex(color.RGB{255, 255, 255}) //white
	if hex == "ffffff" {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHex()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHex()! Returned %v instead of ffffff.", hex)
	}

	hex = color.RGBtoHex(color.RGB{0, 128, 0}) //green
	if hex == "008000" {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHex()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHex()! Returned %v instead of 008000", hex)
	}
}

func TestRGBtoDecimal(t *testing.T) {
	var decimal color.Decimal

	decimal = color.RGBtoDecimal(color.RGB{0, 0, 0}) //black
	if decimal == 0 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoDecimal()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoDecimal()! Returned %v instead of 000000.", decimal)
	}

	decimal = color.RGBtoDecimal(color.RGB{255, 255, 255}) //white
	if decimal == 16777215 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHex()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHex()! Returned %v instead of ffffff.", decimal)
	}

	decimal = color.RGBtoDecimal(color.RGB{0, 128, 0}) //green
	if decimal == 32768 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHex()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHex()! Returned %v instead of 008000", decimal)
	}
}

func TestRGBtoCMYK(t *testing.T) {
	var cmyk color.CMYK

	cmyk = color.RGBtoCMYK(color.RGB{0, 0, 0}) //black
	if cmyk.C == 0 && cmyk.M == 0 && cmyk.Y == 0 && cmyk.K == 100 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoCMYK()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoCMYK()! Returned %v, %v, %v, %v instead of 0, 0, 0, 100.", cmyk.C, cmyk.M, cmyk.Y, cmyk.K)
	}

	cmyk = color.RGBtoCMYK(color.RGB{255, 255, 255}) //white
	if cmyk.C == 0 && cmyk.M == 0 && cmyk.Y == 0 && cmyk.K == 0 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoCMYK()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoCMYK()! Returned %v, %v, %v, %v instead of 0, 0, 0, 0.", cmyk.C, cmyk.M, cmyk.Y, cmyk.K)
	}

	cmyk = color.RGBtoCMYK(color.RGB{0, 128, 0}) //green
	if cmyk.C == 100 && cmyk.M == 0 && cmyk.Y == 100 && cmyk.K == 50 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoCMYK()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoCMYK()! Returned %v, %v, %v, %v instead of 100, 0, 100, 50.", cmyk.C, cmyk.M, cmyk.Y, cmyk.K)
	}
}

//Test HSVtoHSL and HSLtoHSV
func TestHSVtoHSL(t *testing.T) {
	var hsl color.HSL

	hsl = color.HSVtoHSL(color.HSV{0, 0, 0}) //black
	if hsl.H == 0 && hsl.S == 0 && hsl.L == 0 {
		fmt.Println("# Black (HSV: 0, 0, 0) is working with HSVtoHSL()")
	} else {
		t.Errorf("Black (HSV: 0, 0, 0) has a problem with HSVtoHSL()! Returned %v, %v, %v instead of 0, 0, 0.", hsl.H, hsl.S, hsl.L)
	}

	hsl = color.HSVtoHSL(color.HSV{0, 0, 100}) //white
	if hsl.H == 0 && hsl.S == 0 && hsl.L == 100 {
		fmt.Println("# White (RGB: 0, 0, 100) is working with HSVtoHSL()")
	} else {
		t.Errorf("White (RGB: 0, 0, 100) has a problem with HSVtoHSL()! Returned %v, %v, %v instead of 0, 0, 100.", hsl.H, hsl.S, hsl.L)
	}

	hsl = color.HSVtoHSL(color.HSV{120, 100, 50}) //green
	if hsl.H == 120 && hsl.S == 100 && hsl.L == 25 {
		fmt.Println("# Green (HSV: 120, 100, 50) is working with HSVtoHSL()")
	} else {
		t.Errorf("Green (HSV: 120, 100, 50) has a problem with HSVtoHSL()! Returned %v, %v, %v instead of 120, 100, 50.", hsl.H, hsl.S, hsl.L)
	}
}

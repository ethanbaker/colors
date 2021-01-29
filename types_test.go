package color_test

import (
	"fmt"
	"testing"

	"github.com/ethanbaker/colors"
)

//Test TYPEtoRGB
func TestHSVtoRGB(t *testing.T) {
	var r, g, b int

	r, g, b = color.HSVtoRGB(0, 0, 0) //black
	if r == 0 && g == 0 && b == 0 {
		fmt.Println("# Black (HSV: 0, 0, 0) is working with HSVtoRGB()")
	} else {
		t.Errorf("Black (HSV: 0, 0, 0) not working with HSVtoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", r, b, g)
	}

	r, g, b = color.HSVtoRGB(0, 0, 100) //white
	if r == 255 && g == 255 && b == 255 {
		fmt.Println("# White (HSV: 0, 0, 100) is working with HSVtoRGB()")
	} else {
		t.Errorf("White (HSV: 0, 0, 100) not working with HSVtoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", r, g, b)
	}

	r, g, b = color.HSVtoRGB(120, 100, 50) //green
	if r == 0 && g == 128 && b == 0 {
		fmt.Println("# Green (HSV: 120, 100, 50) is working with HSVtoRGB()")
	} else {
		t.Errorf("Green (HSV: 120, 100, 50) not working with HSVtoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", r, g, b)
	}
}

func TestHSLtoRGB(t *testing.T) {
	var r, g, b int

	r, g, b = color.HSLtoRGB(0, 0, 0) //black
	if r == 0 && g == 0 && b == 0 {
		fmt.Println("# Black (HSL: 0, 0, 0) is working with HSLtoRGB()")
	} else {
		t.Errorf("Black (HSL: 0, 0, 0) has a problem with HSLtoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", r, g, b)
	}

	r, g, b = color.HSLtoRGB(0, 0, 100) //white
	if r == 255 && g == 255 && b == 255 {
		fmt.Println("# White (HSL: 0, 0, 100) is working with HSLtoRGB()")
	} else {
		t.Errorf("White (HSL: 0, 0, 100) has a problem with HSLtoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", r, g, b)
	}

	r, g, b = color.HSLtoRGB(120, 100, 25) //green
	if r == 0 && g == 128 && b == 0 {
		fmt.Println("# Green (HSL: 120, 100, 25) is working with HSLtoRGB()")
	} else {
		t.Errorf("Green (HSL: 120, 100, 25) has a problem with HSLtoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", r, g, b)
	}
}

func TestHexNametoRGB(t *testing.T) {
	var r, g, b int

	r, g, b = color.HexNametoRGB("#000000") //black
	if r == 0 && g == 0 && b == 0 {
		fmt.Println("# Black (Hex: #000000) is working with HextoRGB()")
	} else {
		t.Errorf("Black (Hex: #000000) has a problem with HextoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", r, g, b)
	}

	r, g, b = color.HexNametoRGB("#ffffff") //white
	if r == 255 && g == 255 && b == 255 {
		fmt.Println("# White (Hex: #ffffff) is working with HextoRGB()")
	} else {
		t.Errorf("White (Hex: #ffffff) has a problem with HextoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", r, g, b)
	}

	r, g, b = color.HexNametoRGB("008000") //green
	if r == 0 && g == 128 && b == 0 {
		fmt.Println("# Green (Hex: #008000) is working with HextoRGB()")
	} else {
		t.Errorf("Green (Hex: #008000) has a problem with HextoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", r, g, b)
	}
}

func TestHexValtoRGB(t *testing.T) {
	var r, g, b int

	r, g, b = color.HexValtoRGB(0) //black
	if r == 0 && g == 0 && b == 0 {
		fmt.Println("# Black (Hex: #000000) is working with HextoRGB()")
	} else {
		t.Errorf("Black (Hex: #000000) has a problem with HextoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", r, g, b)
	}

	r, g, b = color.HexValtoRGB(16777215) //white
	if r == 255 && g == 255 && b == 255 {
		fmt.Println("# White (Hex: #ffffff) is working with HextoRGB()")
	} else {
		t.Errorf("White (Hex: #ffffff) has a problem with HextoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", r, g, b)
	}

	r, g, b = color.HexValtoRGB(32768) //green
	if r == 0 && g == 128 && b == 0 {
		fmt.Println("# Green (Hex: #008000) is working with HextoRGB()")
	} else {
		t.Errorf("Green (Hex: #008000) has a problem with HextoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", r, g, b)
	}
}

func TestCMYKtoRGB(t *testing.T) {
	var r, g, b int

	r, g, b = color.CMYKtoRGB(0, 0, 0, 100) //black
	if r == 0 && g == 0 && b == 0 {
		fmt.Println("# Black (CMYK: 0, 0, 0, 100) is working with CMYKtoRGB()")
	} else {
		t.Errorf("Black (CMYK: 0, 0, 0, 100) has a problem with CMYKtoRGB()! Returned %v, %v, %v instead of 0, 0, 0.", r, g, b)
	}

	r, g, b = color.CMYKtoRGB(0, 0, 0, 0) //white
	if r == 255 && g == 255 && b == 255 {
		fmt.Println("# White (CMYK: 0, 0, 0, 0) is working with CMYKtoRGB()")
	} else {
		t.Errorf("White (CMYK: 0, 0, 0, 100) has a problem with CMYKtoRGB()! Returned %v, %v, %v instead of 255, 255, 255.", r, g, b)
	}

	r, g, b = color.CMYKtoRGB(100, 0, 100, 50) //green
	if r == 0 && g == 128 && b == 0 {
		fmt.Println("# Green (CMYK: 100, 0, 100, 50) is working with CMYKtoRGB()")
	} else {
		t.Errorf("Green (CMYK: 100, 0, 100, 50) has a problem with CMYKtoRGB()! Returned %v, %v, %v instead of 0, 128, 0.", r, g, b)
	}
}

//Test RGBtoTYPE
func TestRGBtoHSV(t *testing.T) {
	var h, s, v int

	h, s, v = color.RGBtoHSV(0, 0, 0) //black
	if h == 0 && s == 0 && v == 0 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoHSV()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoHSV()! Returned %v, %v, %v instead of 0, 0, 0.", h, s, v)
	}

	h, s, v = color.RGBtoHSV(255, 255, 255) //white
	if h == 0 && s == 0 && v == 100 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHSV()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHSV()! Returned %v, %v, %v instead of 0, 0, 100.", h, s, v)
	}

	h, s, v = color.RGBtoHSV(0, 128, 0) //green
	if h == 120 && s == 100 && v == 50 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHSV()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHSV()! Returned %v, %v, %v instead of 120, 100, 50.", h, s, v)
	}
}

func TestRGBtoHSL(t *testing.T) {
	var h, s, l int

	h, s, l = color.RGBtoHSL(0, 0, 0) //black
	if h == 0 && s == 0 && l == 0 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoHSL()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoHSL()! Returned %v, %v, %v instead of 0, 0, 0.", h, s, l)
	}

	h, s, l = color.RGBtoHSL(255, 255, 255) //white
	if h == 0 && s == 0 && l == 100 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHSL()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHSL()! Returned %v, %v, %v instead of 0, 0, 100.", h, s, l)
	}

	h, s, l = color.RGBtoHSL(0, 128, 0) //green
	if h == 120 && s == 100 && l == 25 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHSL()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHSL()! Returned %v, %v, %v instead of 120, 100, 25.", h, s, l)
	}
}

func TestRGBtoHexName(t *testing.T) {
	var hex string

	hex = color.RGBtoHexName(0, 0, 0) //black
	if hex == "000000" {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoHex()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoHex()! Returned %v instead of 000000.", hex)
	}

	hex = color.RGBtoHexName(255, 255, 255) //white
	if hex == "ffffff" {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHex()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHex()! Returned %v instead of ffffff.", hex)
	}

	hex = color.RGBtoHexName(0, 128, 0) //green
	if hex == "008000" {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHex()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHex()! Returned %v instead of 008000", hex)
	}
}

func TestRGBtoHexVal(t *testing.T) {
	var hex int

	hex = color.RGBtoHexVal(0, 0, 0) //black
	if hex == 0 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoHex()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoHex()! Returned %v instead of 000000.", hex)
	}

	hex = color.RGBtoHexVal(255, 255, 255) //white
	if hex == 16777215 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoHex()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoHex()! Returned %v instead of ffffff.", hex)
	}

	hex = color.RGBtoHexVal(0, 128, 0) //green
	if hex == 32768 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoHex()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoHex()! Returned %v instead of 008000", hex)
	}
}

func TestRGBtoCMYK(t *testing.T) {
	var c, m, y, k int

	c, m, y, k = color.RGBtoCMYK(0, 0, 0) //black
	if c == 0 && m == 0 && y == 0 && k == 100 {
		fmt.Println("# Black (RGB: 0, 0, 0) is working with RGBtoCMYK()")
	} else {
		t.Errorf("Black (RGB: 0, 0, 0) has a problem with RGBtoCMYK()! Returned %v, %v, %v, %v instead of 0, 0, 0, 100.", c, m, y, k)
	}

	c, m, y, k = color.RGBtoCMYK(255, 255, 255) //white
	if c == 0 && m == 0 && y == 0 && k == 0 {
		fmt.Println("# White (RGB: 255, 255, 255) is working with RGBtoCMYK()")
	} else {
		t.Errorf("White (RGB: 255, 255, 255) has a problem with RGBtoCMYK()! Returned %v, %v, %v, %v instead of 0, 0, 0, 0.", c, m, y, k)
	}

	c, m, y, k = color.RGBtoCMYK(0, 128, 0) //green
	if c == 100 && m == 0 && y == 100 && k == 50 {
		fmt.Println("# Green (RGB: 0, 128, 0) is working with RGBtoCMYK()")
	} else {
		t.Errorf("Green (RGB: 0, 128, 0) has a problem with RGBtoCMYK()! Returned %v, %v, %v, %v instead of 100, 0, 100, 50.", c, m, y, k)
	}
}

//Test HSVtoHSL and HSLtoHSV
func TestHSVtoHSL(t *testing.T) {
	var h, s, l int

	h, s, l = color.RGBtoHSL(0, 0, 0) //black
	if h == 0 && s == 0 && l == 0 {
		fmt.Println("# Black (HSV: 0, 0, 0) is working with HSVtoHSL()")
	} else {
		t.Errorf("Black (HSV: 0, 0, 0) has a problem with HSVtoHSL()! Returned %v, %v, %v instead of 0, 0, 0.", h, s, l)
	}

	h, s, l = color.HSVtoHSL(0, 0, 100) //white
	if h == 0 && s == 0 && l == 100 {
		fmt.Println("# White (RGB: 0, 0, 100) is working with HSVtoHSL()")
	} else {
		t.Errorf("White (RGB: 0, 0, 100) has a problem with HSVtoHSL()! Returned %v, %v, %v instead of 0, 0, 100.", h, s, l)
	}

	h, s, l = color.HSVtoHSL(120, 100, 50) //green
	if h == 120 && s == 100 && l == 25 {
		fmt.Println("# Green (HSV: 120, 100, 50) is working with HSVtoHSL()")
	} else {
		t.Errorf("Green (HSV: 120, 100, 50) has a problem with HSVtoHSL()! Returned %v, %v, %v instead of 120, 100, 50.", h, s, l)
	}
}

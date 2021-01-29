package color_test

import (
	"fmt"
	"testing"

	"gitlab.com/ethanbakerdev/colors"
	"gitlab.com/ethanbakerdev/colors/css"
	"gitlab.com/ethanbakerdev/colors/sol"
)

func TestDecolorAnsiSol(t *testing.T) {
	s := sol.RandomAnsi() + "random" + sol.AnsiX
	t.Logf("# len(%v) is %v \n", s, len(s))
	s = color.DecolorAnsiSol(s)
	t.Logf("len(%v) is %v \n", s, len(s))
	s = color.DecolorAnsiSol(s)
	if len(s) != 10 {
		t.Errorf("%v has problems with DecolorAnsiSol", s)
	}
}

func TestSearchAnsiIndex(t *testing.T) {
	c := color.DecimaltoAnsi(css.Cyan) + "css.cyan"
	sc := color.SearchAnsiIndex("css.cyan")
	s := color.DecimaltoAnsi(sol.Cyan) + "sol.cyan"
	ss := color.SearchAnsiIndex("sol.cyan")
	if sc != c {
		t.Errorf("css.cyan has problems with SearchAnsiIndex! Returned: %v %v", c, sol.AnsiReset)
	} else {
		fmt.Printf("# %v found SearchAnsiIndex!%v\n", c, css.AnsiReset)
	}
	if s != ss {
		t.Errorf("sol.cyan has problems with SearchAnsiIndex! Returned: %v %v", s, sol.AnsiReset)
	} else {
		fmt.Printf("# %v found in SearchAnsiIndex!%v\n", s, sol.AnsiReset)
	}
}

func TestSearchHexIndex(t *testing.T) {
	var c color.Decimal = 0x00ffff
	sc := color.SearchHexIndex("css.cyan")
	var s color.Decimal = 0x2aa198
	ss := color.SearchHexIndex("sol.cyan")
	if sc != c {
		t.Errorf("%v%v (css.cyan) has problems with SearchHexIndex", css.Cyan, sc)
	} else {
		fmt.Printf("# %v%v (css.cyan) found SearchHexIndex!%v\n", css.Cyan, c, css.AnsiReset)
	}
	if s != ss {
		t.Errorf("%v%v (sol.cyan) has problems with SearchHexIndex", sol.Cyan, s)
	} else {
		fmt.Printf("# %v%v (sol.cyan) found in SearchHexIndex!%v\n", sol.Cyan, s, sol.AnsiReset)
	}
}

func TestRGB(t *testing.T) {
	s := color.AnsiRGB(color.RGB{0, 255, 255}) + "cyan"
	if len(s) != 21 {
		t.Errorf("AnsiRgb() is not working with parameter color.RGB{0, 255, 255}! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with AnsiRgb() %v\n", s, css.AnsiReset)
	}

}

func TestRGBBackground(t *testing.T) {
	s := color.AnsiRGBBackground(color.RGB{0, 255, 255}) + "cyan background"
	if len(s) != 32 {
		t.Errorf("AnsiRgbBackground() is not working with parameter color.RGB{0, 255, 255}! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with AnsiRgbBackground() %v\n", s, css.AnsiReset)
	}
}

/*
func TestHexRgb(t *testing.T) {
	v := color.HexRgb(0, 255, 255)
	if v != 65535 {
		t.Errorf("HexRgb() is not working with parameters 0, 255, and 255! Output: %v", v)
	} else {
		fmt.Printf("#%v %v made with HexRgb() %v\n", css.Cyan, v, css.AnsiReset)
	}
}

func TestHexRgbName(t *testing.T) {
	s := color.HexRgbName(0, 255, 255)
	if s != "#00ffff" {
		t.Errorf("HexRgbName() is not working with parameters 0, 255, and 255! Output: %v", s)
	} else {
		fmt.Printf("#%v %v made with HexRgbName() %v\n", css.Cyan, s, css.AnsiReset)
	}
}
*/

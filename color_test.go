package color_test

import (
	"fmt"
	"testing"

	"github.com/ethanbaker/colors"
	"github.com/ethanbaker/colors/css"
	"github.com/ethanbaker/colors/sol"
)

func TestDecolorAsciiSol(t *testing.T) {
	s := sol.RandomAscii() + "random" + sol.AsciiX
	t.Logf("# len(%v) is %v \n", s, len(s))
	s = color.DecolorAsciiSol(s)
	t.Logf("len(%v) is %v \n", s, len(s))
	s = color.DecolorAsciiSol(s)
	if len(s) != 10 {
		t.Errorf("%v has problems with DecolorAsciiSol", s)
	}
}

func TestSearchAsciiIndex(t *testing.T) {
	c := color.HexValtoAscii(css.Cyan) + "css.cyan"
	sc := color.SearchAsciiIndex("css.cyan")
	s := color.HexValtoAscii(sol.Cyan) + "sol.cyan"
	ss := color.SearchAsciiIndex("sol.cyan")
	if sc != c {
		t.Errorf("css.cyan has problems with SearchAsciiIndex! Returned: %v %v", c, sol.AsciiReset)
	} else {
		fmt.Printf("# %v found SearchAsciiIndex!%v\n", c, css.AsciiReset)
	}
	if s != ss {
		t.Errorf("sol.cyan has problems with SearchAsciiIndex! Returned: %v %v", s, sol.AsciiReset)
	} else {
		fmt.Printf("# %v found in SearchAsciiIndex!%v\n", s, sol.AsciiReset)
	}
}

func TestSearchHexIndex(t *testing.T) {
	var c int32 = 0x00ffff
	sc := color.SearchHexIndex("css.cyan")
	var s int32 = 0x2aa198
	ss := color.SearchHexIndex("sol.cyan")
	if sc != c {
		t.Errorf("%v%v (css.cyan) has problems with SearchHexIndex", css.Cyan, sc)
	} else {
		fmt.Printf("# %v%v (css.cyan) found SearchHexIndex!%v\n", css.Cyan, c, css.AsciiReset)
	}
	if s != ss {
		t.Errorf("%v%v (sol.cyan) has problems with SearchHexIndex", sol.Cyan, s)
	} else {
		fmt.Printf("# %v%v (sol.cyan) found in SearchHexIndex!%v\n", sol.Cyan, s, sol.AsciiReset)
	}
}

func TestRgb(t *testing.T) {
	s := color.AsciiRgb("0", "255", "255") + "cyan"
	if len(s) != 21 {
		t.Errorf("AsciiRgb() is not working with paramaters \"0\", \"255\", and \"255\"! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with AsciiRgb() %v\n", s, css.AsciiReset)
	}

}

func TestRgbBackground(t *testing.T) {
	s := color.AsciiRgbBackground("0", "255", "255") + "cyan background"
	if len(s) != 32 {
		t.Errorf("AsciiRgbBackground() is not working with parameters \"0\", \"255\", and \"255\"! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with AsciiRgbBackground() %v\n", s, css.AsciiReset)
	}
}

/*
func TestHexRgb(t *testing.T) {
	v := color.HexRgb(0, 255, 255)
	if v != 65535 {
		t.Errorf("HexRgb() is not working with parameters 0, 255, and 255! Output: %v", v)
	} else {
		fmt.Printf("#%v %v made with HexRgb() %v\n", css.Cyan, v, css.AsciiReset)
	}
}

func TestHexRgbName(t *testing.T) {
	s := color.HexRgbName(0, 255, 255)
	if s != "#00ffff" {
		t.Errorf("HexRgbName() is not working with parameters 0, 255, and 255! Output: %v", s)
	} else {
		fmt.Printf("#%v %v made with HexRgbName() %v\n", css.Cyan, s, css.AsciiReset)
	}
}
*/

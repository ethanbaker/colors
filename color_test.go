package color_test

import (
	"fmt"
	"testing"

	"github.com/ethanbaker/go-colors"
	"github.com/ethanbaker/go-colors/css"
	"github.com/ethanbaker/go-colors/sol"
)

func TestDecolorAsciiSol(t *testing.T) {
	s := sol.Random() + "random" + sol.X
	t.Logf("# len(%v) is %v \n", s, len(s))
	s = color.DecolorAsciiSol(s)
	t.Logf("len(%v) is %v \n", s, len(s))
	s = color.DecolorAsciiSol(s)
	if len(s) != 6 {
		t.Errorf("%v has problems with DecolorAsciiSol", s)
	}
}

func TestSearchAsciiIndex(t *testing.T) {
	c := "\x1b[38;2;0;255;255mcss.cyan"
	sc := color.SearchAsciiIndex("css.cyan")
	s := "\033[0;36msol.cyan"
	ss := color.SearchAsciiIndex("sol.cyan")
	if sc != c {
		t.Errorf("%v has problems with SearchAsciiIndex", c)
	} else {
		fmt.Printf("%v# %v found SearchAsciiIndex!%v\n", css.Reset, c, css.Reset)
	}
	if s != ss {
		t.Errorf("%v has problems with SearchAsciiIndex", s)
	} else {
		fmt.Printf("# %v found in SearchAsciiIndex!%v\n", s, sol.Reset)
	}
}

func TestRgb(t *testing.T) {
	s := color.AsciiRgb("0", "255", "255") + "cyan"
	if len(s) != 21 {
		t.Errorf("AsciiRgb() is not working with paramaters \"0\", \"255\", and \"255\"! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with AsciiRgb() %v\n", s, css.Reset)
	}

}

func TestRgbBackground(t *testing.T) {
	s := color.AsciiRgbBackground("0", "255", "255") + "cyan background"
	if len(s) != 32 {
		t.Errorf("AsciiRgbBackground() is not working with parameters \"0\", \"255\", and \"255\"! Output: %v", len(s))
	} else {
		fmt.Printf("#%v %v made with AsciiRgbBackground() %v\n", css.Black, s, css.Reset)
	}
}

func TestSearchHexIndex(t *testing.T) {
	c := 0x00ffff
	sc := color.SearchHexIndex("css.cyan")
	s := 0x2aa198
	ss := color.SearchHexIndex("sol.cyan")
	if sc != c {
		t.Errorf("%v%v (css.cyan) has problems with SearchHexIndex", css.Cyan, sc)
	} else {
		fmt.Printf("%v# %v%v (css.cyan) found SearchHexIndex!%v\n", css.Reset, css.Cyan, c, css.Reset)
	}
	if s != ss {
		t.Errorf("%v%v (sol.cyan) has problems with SearchHexIndex", sol.Cyan, s)
	} else {
		fmt.Printf("# %v%v (sol.cyan) found in SearchHexIndex!%v\n", sol.Cyan, s, sol.Reset)
	}
}

func TestHexRgb(t *testing.T) {
	v := color.HexRgb(0, 255, 255)
	if v != 65535 {
		t.Errorf("HexRgb() is not working with parameters 0, 255, and 255! Output: %v", v)
	} else {
		fmt.Printf("#%v %v made with HexRgb() %v\n", css.Cyan, v, css.Reset)
	}
}

func TestHexRgbName(t *testing.T) {
	s := color.HexRgbName(0, 255, 255)
	if s != "#00ffff" {
		t.Errorf("HexRgbName() is not working with parameters 0, 255, and 255! Output: %v", s)
	} else {
		fmt.Printf("#%v %v made with HexRgbName() %v\n", css.Cyan, s, css.Reset)
	}
}

package color_test

import (
	"fmt"
	"testing"

	"github.com/ethanbaker/go-colors"
	"github.com/ethanbaker/go-colors/css"
	"github.com/ethanbaker/go-colors/sol"
)

func TestDecolorSol(t *testing.T) {
	s := sol.Random() + "random" + sol.X
	t.Logf("# len(%v) is %v \n", s, len(s))
	s = color.DecolorSol(s)
	t.Logf("len(%v) is %v \n", s, len(s))
	s = color.DecolorSol(s)
	if len(s) != 6 {
		t.Errorf("%v has problems with DecolorSol", s)
	}
}

func TestSearchGlobalIndex(t *testing.T) {
	c := "\x1b[38;2;0;255;255mcss.cyan"
	sc := color.SearchGlobalIndex("css.cyan")
	s := "\033[0;36msol.cyan"
	ss := color.SearchGlobalIndex("sol.cyan")
	if sc != c {
		t.Errorf("%v has problems with SearchGlobalIndex", c)
	} else {
		fmt.Printf("%v# %v found SearchGlobalIndex!%v\n", css.Reset, c, css.Reset)
	}
	if s != ss {
		t.Errorf("%v has problems with SearchGlobalIndex", s)
	} else {
		fmt.Printf("# %v found in SearchGlobalIndex!%v\n", s, sol.Reset)
	}
}

func TestRgb(t *testing.T) {
	s := color.Rgb("0", "255", "255") + "cyan"
	if len(s) != 21 {
		t.Errorf("Rgb() is not working with paramaters \"0\", \"255\", and \"255\"! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with Rgb() %v\n", s, css.Reset)
	}

}

func TestRgbBackground(t *testing.T) {
	s := color.RgbBackground("0", "255", "255") + "cyan background"
	if len(s) != 32 {
		t.Errorf("RgbBackground() is not working with parameters \"0\", \"255\", and \"255\"! Output: %v", len(s))
	} else {
		fmt.Printf("#%v %v made with RgbBackground() %v\n", css.Black, s, css.Reset)
	}
}

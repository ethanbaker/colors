package color_test

import (
	"fmt"
	"testing"

	color "github.com/ethanbaker/colors"
	"github.com/ethanbaker/colors/css"
	"github.com/ethanbaker/colors/sol"
)

func TestDecolorAnsiSol(t *testing.T) {
	s := sol.Random() + "random" + sol.X
	t.Logf("# len(%v) is %v \n", s, len(s))
	s = color.DecolorAnsiSol(s)
	t.Logf("len(%v) is %v \n", s, len(s))
	s = color.DecolorAnsiSol(s)
	if len(s) != 10 {
		t.Errorf("%v has problems with DecolorAnsiSol", s)
	}
}

func TestDecolorAnsiCss(t *testing.T) {
	s := css.Random() + "random" + sol.X
	t.Logf("# len(%v) is %v \n", s, len(s))
	s = color.DecolorAnsiCss(s)
	t.Logf("len(%v) is %v \n", s, len(s))
	s = color.DecolorAnsiSol(s)
	if len(s) != 10 {
		t.Errorf("%v has problems with DecolorAnsiCss", s)
	}
}

func TestRGB(t *testing.T) {
	s := color.AnsiRGB(color.RGB{R: 0, G: 255, B: 255}) + "cyan"
	if len(s) != 21 {
		t.Errorf("AnsiRgb() is not working with parameter color.RGB{0, 255, 255}! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with AnsiRgb() %v\n", s, css.Reset)
	}
}

func TestRGBBackground(t *testing.T) {
	s := color.AnsiRGBBackground(color.RGB{R: 0, G: 255, B: 255}) + "cyan background"
	if len(s) != 32 {
		t.Errorf("AnsiRgbBackground() is not working with parameter color.RGB{0, 255, 255}! Output: %v", len(s))
	} else {
		fmt.Printf("# %v made with AnsiRgbBackground() %v\n", s, css.Reset)
	}
}

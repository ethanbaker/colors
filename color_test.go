package color_test

import (
	"fmt"
	"testing"

	"gitlab.com/ethanbaker.dev/color"
	"gitlab.com/ethanbaker.dev/color/sol"
        "gitlab.com/ethanbaker.dev/color/css"
)

func TestDecolor(t *testing.T) {
	s := sol.Random() + "random" + sol.X
	fmt.Printf("# len(%v) is %v \n", s, len(s))
	s = color.Decolor(s)
	fmt.Printf("# len(%v) is %v \n", s, len(s))
	s = color.Decolor(s)
	if len(s) != 6 {
		t.Errorf("%v has problems with Decolor", s)
	}
}

func TestStripSol(t *testing.T) {
	s := sol.Random() + "random" + sol.X + sol.ClearScreen
	s = color.Strip(s)
	if len(s) != 6 {
		t.Errorf("%v has problems with Strip", s)
	}
}

func TestStripCss(t *testing.T) {
        c := sol.Random() + "random" + sol.X + sol.ClearScreen
        c = color.Strip(c)
        if len(s) != 6 {
                t.Errorf("%v has problems with Strip", c)
        }
}

package color_test

import (
	"testing"

	"gitlab.com/skilstak/code/go/color"
	"gitlab.com/skilstak/code/go/color/sol"
)

func TestDecolor(t *testing.T) {
	s := sol.Random() + "random" + sol.X
	t.Logf("# len(%v) is %v \n", s, len(s))
	s = color.Decolor(s)
	t.Logf("len(%v) is %v \n", s, len(s))
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
	if len(c) != 6 {
		t.Errorf("%v has problems with Strip", c)
	}
}

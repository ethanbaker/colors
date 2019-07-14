package color_test

import (
	"fmt"
	"testing"

	"gitlab.com/ethandbaker/color"
	"gitlab.com/ethandbaker/color/sol"
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

func TestStrip(t *testing.T) {
	s := sol.Random() + "random" + sol.X + sol.ClearScreen
	s = color.Strip(s)
	if len(s) != 6 {
		t.Errorf("%v has problems with Strip", s)
	}
}

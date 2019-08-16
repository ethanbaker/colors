package sol_test

import (
	"fmt"
	"testing"

	"gitlab.com/skilstak/code/go/color/sol"
)

func ExampleCyan() {

	// also see color/cmd/sol command for example

	fmt.Print(sol.Cyan + "Cyan" + sol.Reset + " ")
	fmt.Print(sol.C + "Cyan" + sol.X + " ")
	fmt.Printf("%v%v%v ", sol.C, "Cyan", sol.X)

}

func TestRandom(t *testing.T) {
	s := sol.Random() + "random" + sol.X
	fmt.Printf("# len(%v) is %v \n", s, len(s))
	if len(s) != 17 {
		t.Errorf("%v has problems with Random", s)
	}
}

func TestMultiple(t *testing.T) {
	s := sol.Multiple("multiple")
	fmt.Printf("# len(%v) is %v \n", s, len(s))
	if len(s) != 68 {
		t.Errorf("%v has problems with Multiple", s)
	}
}

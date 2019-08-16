package css_test

import (
        "fmt"
        "testing"

	"gitlab.com/skilstak/code/go/color/css"
)

func TestRandom(t *testing.T) {
        c := css.Random() + "random" + css.X
        fmt.Printf(css.X + "# len(%v) is %v \n", c, len(c))
        if false { //TODO Fix if statement
                t.Errorf("%v has problems with Random", c) 
        }
}

func TestColors(t *testing.T) {
        i := 0
        for range css.Index {
                i += 1
        }
        if i != 141 {
                t.Errorf("One or more of the colors are not working.")
        }
}

func ExampleCyan() {
        fmt.Print(css.Cyan + "Cyan" + css.X + " ")
        fmt.Printf("%v%v%v ", css.Cyan, "Cyan", css.X)
}

package css_test

import (
	"fmt"
        "testing"

	"gitlab.com/skilstak/code/go/color/css"
)

func TestAll(t *testing.T) {
        for k := range css.Index {
                    fmt.Printf("%s%s\n", css.Index[k], k)
        }
}


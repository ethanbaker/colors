package css_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethanbaker/colors/css"
	"github.com/stretchr/testify/assert"
)

// Constant string to modify in different functions
const multipleString = "MULTIPLE"

// Print an example use case of using the color cyan to print text
func ExampleCyan() {
	fmt.Println(css.Cyan + "Cyan" + css.Reset)
	fmt.Printf("%v%v%v\n", css.Cyan, "Cyan", css.Reset)
}

func TestRandom(t *testing.T) {
	assert := assert.New(t)

	s := css.Random() + "random" + css.X
	t.Logf("Random string returned: %v", s)

	matches := false
	for _, v := range css.AnsiMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random color should be a cssarized color")
}

func TestRandomBackground(t *testing.T) {
	assert := assert.New(t)

	s := css.RandomBackground() + "random" + css.X
	t.Logf("Random background string returned: %v", s)

	matches := false
	for _, v := range css.AnsiBackgroundMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random background color should be a cssarized background")
}

func TestRandomHex(t *testing.T) {
	assert := assert.New(t)

	n := css.RandomHex()
	t.Logf("Random hex color returned: %T", n)
	matches := false
	for _, v := range css.HexMap {
		if v == n {
			matches = true
			break
		}
	}
	assert.True(matches, "Random hex color should be a cssarized hex")
}

func TestMultiple(t *testing.T) {
	assert := assert.New(t)

	s := css.Multiple(multipleString)
	t.Logf("Multiple string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range css.AnsiMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a cssarized color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

func TestMultipleBackground(t *testing.T) {
	assert := assert.New(t)

	s := css.MultipleBackground("MULTIPLE")
	t.Logf("Multiple background string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range css.AnsiBackgroundMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a cssarized background color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

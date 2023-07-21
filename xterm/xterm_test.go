package xterm_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethanbaker/colors/xterm"
	"github.com/stretchr/testify/assert"
)

// Constant string to modify in different functions
const multipleString = "MULTIPLE"

// Print an example use case of using the color cyan to print text
func ExampleCyan1() {
	fmt.Println(xterm.Cyan1 + "Cyan" + xterm.Reset)
	fmt.Printf("%v%v%v\n", xterm.Cyan1, "Cyan", xterm.Reset)
}

func TestRandom(t *testing.T) {
	assert := assert.New(t)

	s := xterm.Random() + "random" + xterm.X
	t.Logf("Random string returned: %v", s)

	matches := false
	for _, v := range xterm.AnsiMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random color should be a xtermarized color")
}

func TestRandomBackground(t *testing.T) {
	assert := assert.New(t)

	s := xterm.RandomBackground() + "random" + xterm.X
	t.Logf("Random background string returned: %v", s)

	matches := false
	for _, v := range xterm.AnsiBackgroundMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random background color should be a xtermarized background")
}

func TestRandomHex(t *testing.T) {
	assert := assert.New(t)

	n := xterm.RandomHex()
	t.Logf("Random hex color returned: %T", n)
	matches := false
	for _, v := range xterm.HexMap {
		if v == n {
			matches = true
			break
		}
	}
	assert.True(matches, "Random hex color should be a xtermarized hex")
}

func TestMultiple(t *testing.T) {
	assert := assert.New(t)

	s := xterm.Multiple(multipleString)
	t.Logf("Multiple string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range xterm.AnsiMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a xtermarized color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

func TestMultipleBackground(t *testing.T) {
	assert := assert.New(t)

	s := xterm.MultipleBackground("MULTIPLE")
	t.Logf("Multiple background string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range xterm.AnsiBackgroundMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a xtermarized background color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

package sol_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethanbaker/colors/sol"
	"github.com/stretchr/testify/assert"
)

// Constant string to modify in different functions
const multipleString = "MULTIPLE"

// Print an example use case of using the color cyan to print text
func ExampleCyan() {
	fmt.Println(sol.Cyan + "Cyan" + sol.Reset)
	fmt.Println(sol.C + "Cyan" + sol.X)
	fmt.Printf("%v%v%v\n", sol.C, "Cyan", sol.X)
}

func TestRandom(t *testing.T) {
	assert := assert.New(t)

	s := sol.Random() + "random" + sol.X
	t.Logf("Random string returned: %v", s)

	matches := false
	for _, v := range sol.AnsiMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random color should be a solarized color")
}

func TestRandomBackground(t *testing.T) {
	assert := assert.New(t)

	s := sol.RandomBackground() + "random" + sol.X
	t.Logf("Random background string returned: %v", s)

	matches := false
	for _, v := range sol.AnsiBackgroundMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random background color should be a solarized background")
}

func TestRandomBase(t *testing.T) {
	assert := assert.New(t)

	s := sol.RandomBase() + "random" + sol.X
	t.Logf("Random background string returned: %v", s)

	matches := false
	for _, v := range sol.AnsiBaseMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random base color should be a solarized base")
}

func TestRandomBaseBackground(t *testing.T) {
	assert := assert.New(t)

	s := sol.RandomBaseBackground() + "random" + sol.X
	t.Logf("Random background string returned: %v", s)

	matches := false
	for _, v := range sol.AnsiBaseBackgroundMap {
		if strings.HasPrefix(s, v) {
			matches = true
			break
		}
	}
	assert.True(matches, "Random base background color should be a solarized base")
}

func TestRandomHex(t *testing.T) {
	assert := assert.New(t)

	n := sol.RandomHex()
	t.Logf("Random hex color returned: %T", n)
	matches := false
	for _, v := range sol.HexMap {
		if v == n {
			matches = true
			break
		}
	}
	assert.True(matches, "Random hex color should be a solarized hex")
}

func TestRandomHexBase(t *testing.T) {
	assert := assert.New(t)

	n := sol.RandomHexBase()
	t.Logf("Random hex color returned: %T", n)
	matches := false
	for _, v := range sol.HexBaseMap {
		if v == n {
			matches = true
			break
		}
	}
	assert.True(matches, "Random hex color should be a solarized hex")
}

func TestMultiple(t *testing.T) {
	assert := assert.New(t)

	s := sol.Multiple(multipleString)
	t.Logf("Multiple string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range sol.AnsiMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a solarized color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

func TestMultipleBase(t *testing.T) {
	assert := assert.New(t)

	s := sol.MultipleBase("MULTIPLE")
	t.Logf("Multiple base string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range sol.AnsiBaseMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a solarized based color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

func TestMultipleBackground(t *testing.T) {
	assert := assert.New(t)

	s := sol.MultipleBackground("MULTIPLE")
	t.Logf("Multiple background string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range sol.AnsiBackgroundMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a solarized background color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

func TestMultipleBaseBackground(t *testing.T) {
	assert := assert.New(t)

	s := sol.MultipleBaseBackground("multiple")
	t.Logf("Multiple base background string returned: %v", s)

	// Loop through each character in the string
	for _, c := range multipleString {
		// Make sure the current character has the right color
		matches := false
		for _, v := range sol.AnsiBaseBackgroundMap {
			if strings.HasPrefix(s, v) {
				matches = true
				break
			}
		}
		assert.True(matches, fmt.Sprintf("Rune '%v' in string '%v' does not have a solarized base background color", string(c), multipleString))

		// Trim to the next rune in the string
		s = s[strings.Index(s, string(c))+1:]
	}
}

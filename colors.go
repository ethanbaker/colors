package colors

import (
	"math/rand"
	"time"
)

const (
	Base03      = "\033[1;30m"
	Base02      = "\033[0;30m"
	Base01      = "\033[1;32m"
	Base00      = "\033[1;33m"
	Base0       = "\033[1;34m"
	Base1       = "\033[1;36m"
	Base2       = "\033[0;37m"
	Base3       = "\033[1;37m"
	Yellow      = "\033[0;33m"
	Orange      = "\033[1;31m"
	Red         = "\033[0;31m"
	Magenta     = "\033[0;35m"
	Violet      = "\033[1;35m"
	Blue        = "\033[0;34m"
	Cyan        = "\033[0;36m"
	Green       = "\033[0;32m"
	Reset       = "\033[0m"
	ClearScreen = "\033[2J\033[H"
	ClearLine   = "\033[2K\033[G"
	CursorOff   = "\033[?25l"
	CursorOn    = "\033[?25h"
)

var SolarizedColors = []string{Yellow, Orange, Red, Magenta, Violet, Blue, Cyan, Green}

func RandomSolarized() string {
	rand.Seed(time.Now().UnixNano())
	return SolarizedColors[rand.Intn(len(SolarizedColors))]
}

func RS() string { return RandomSolarized() }

func Multi(s string) string {
	var m string
	for _, r := range s {
		m += RS() + string(r)
	}
	m += Reset
	return m
}

type Colors struct {
	Off bool
}

func New() Colors {
	c := Colors{}
	return c
}

func (c Colors) Y() string {
	if c.Off {
		return ""
	}
	return Yellow
}

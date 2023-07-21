package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ethanbaker/colors/sol"
)

const usagetxt = `usage: sol COLOR
    COLOR  b02, b01, b00, b0, b1, b2,
           y, o, r, m, v, b, c, g
		   yellow, orange, red, magenta, violet, blue, cyan, green
`

func usage() {
	fmt.Print(usagetxt)
	os.Exit(1)
}

func main() {
	if len(os.Args) <= 2 {
		usage()
	}

	color := strings.ToLower(os.Args[1])

	// If the user wants a sample print it out
	if color == "sample" {
		// Collect a list of sol sample lines
		var lines []string
		for name, val := range sol.AnsiMap {
			lines = append(lines, fmt.Sprintf("%v%v", val, name))
		}

		// Sort the list of lines
		sort.Slice(lines, func(i, j int) bool {
			return lines[i][strings.Index(lines[i], "m")+1:] < lines[j][strings.Index(lines[j], "m")+1:]
		})

		// Print out the list of lines in alphabetical order
		for _, l := range lines {
			fmt.Println(l)
		}

		return
	}

	// Print the color if it exists
	s1, ok1 := sol.AnsiMap[color]
	s2, ok2 := sol.AnsiBaseMap[color]
	if !ok1 && !ok2 {
		usage()
	}

	fmt.Print(s1 + s2)

	// Print out the rest of the command
	for i := 2; i < len(os.Args); i++ {
		fmt.Print(os.Args[i] + " ")
	}
	fmt.Println()
}

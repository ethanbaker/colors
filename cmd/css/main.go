package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ethanbaker/colors/css"
)

const usagetxt = `usage: css COLOR 
-> Print corresponding color 

COLOR: A css color

OR usage: css Sample
-> Shows all color names in the cooresponding color`

func usage() {
	fmt.Println(usagetxt)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	color := strings.ToLower(os.Args[1])

	// If the user wants a sample print it out
	if color == "sample" {
		// Collect a list of css sample lines
		var lines []string
		for name, val := range css.AnsiMap {
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
	val, ok := css.AnsiMap[color]
	if !ok {
		usage()
	}
	fmt.Print(val)

	// Print out the rest of the command
	for i := 2; i < len(os.Args); i++ {
		fmt.Print(os.Args[i] + " ")
	}
	fmt.Println()
}

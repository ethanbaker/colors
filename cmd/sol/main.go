package main

import (
	"fmt"
	"os"

	"gitlab.com/skilstak/code/colors/go"
)

const usagetxt = `
usage: sol COLOR
    COLOR  b02, b01, b00, b0, b1, b2,
           y, o, r, m, v, b, c, g,
           rnd, x, cl, ln, vis, invis
`

func usage() {
	fmt.Println(usagetxt)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	switch os.Args[1] {
	case "b02":
		fmt.Println(colors.SolBase02)
	case "b01":
		fmt.Println(colors.SolBase01)
	case "b00":
		fmt.Println(colors.SolBase00)
	case "b0":
		fmt.Println(colors.SolBase0)
	case "b1":
		fmt.Println(colors.SolBase1)
	case "b2":
		fmt.Println(colors.SolBase2)
	case "y":
		fmt.Println(colors.SolYellow)
	case "o":
		fmt.Println(colors.SolOrange)
	case "r":
		fmt.Println(colors.SolRed)
	case "m":
		fmt.Println(colors.SolMagenta)
	case "v":
		fmt.Println(colors.SolViolet)
	case "b":
		fmt.Println(colors.SolBlue)
	case "c":
		fmt.Println(colors.SolCyan)
	case "g":
		fmt.Println(colors.SolGreen)
	case "rnd":
		fmt.Println(colors.RandSol)
	case "x":
		fmt.Println(colors.Reset)
	case "cl":
		fmt.Println(colors.ClearScreen)
	case "ln":
		fmt.Println(colors.ClearLine)
	case "invis":
		fmt.Println(colors.CursorOff)
	case "vis":
		fmt.Println(colors.CursorOn)
	case "sample":
		fmt.Printf("%v%v ", colors.SolBase02, "SolBase02")
		fmt.Printf("%v%v ", colors.SolBase01, "SolBase01")
		fmt.Printf("%v%v ", colors.SolBase00, "SolBase00")
		fmt.Printf("%v%v ", colors.SolBase0, "SolBase0")
		fmt.Printf("%v%v ", colors.SolBase1, "SolBase1")
		fmt.Printf("%v%v\n", colors.SolBase2, "SolBase2")
		fmt.Printf("%v%v ", colors.SolYellow, "SolYellow")
		fmt.Printf("%v%v ", colors.SolOrange, "SolOrange")
		fmt.Printf("%v%v ", colors.SolRed, "SolRed")
		fmt.Printf("%v%v ", colors.SolMagenta, "SolMagenta")
		fmt.Printf("%v%v ", colors.SolViolet, "SolViolet")
		fmt.Printf("%v%v ", colors.SolBlue, "SolBlue")
		fmt.Printf("%v%v ", colors.SolCyan, "SolCyan")
		fmt.Printf("%v%v\n", colors.SolGreen, "SolGreen")
	default:
		usage()
	}
}

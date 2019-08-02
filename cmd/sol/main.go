package main

import (
	"fmt"
	"os"

	"gitlab.com/ethanbaker.dev/color/sol"
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
		fmt.Println(sol.Base02)
	case "b01":
		fmt.Println(sol.Base01)
	case "b00":
		fmt.Println(sol.Base00)
	case "b0":
		fmt.Println(sol.Base0)
	case "b1":
		fmt.Println(sol.Base1)
	case "b2":
		fmt.Println(sol.Base2)
	case "y":
		fmt.Println(sol.Yellow)
	case "o":
		fmt.Println(sol.Orange)
	case "r":
		fmt.Println(sol.Red)
	case "m":
		fmt.Println(sol.Magenta)
	case "v":
		fmt.Println(sol.Violet)
	case "b":
		fmt.Println(sol.Blue)
	case "c":
		fmt.Println(sol.Cyan)
	case "g":
		fmt.Println(sol.Green)
	case "rnd":
		fmt.Println(sol.Random)
	case "x":
		fmt.Println(sol.Reset)
	case "cl":
		fmt.Println(sol.ClearScreen)
	case "ln":
		fmt.Println(sol.ClearLine)
	case "invis":
		fmt.Println(sol.CursorOff)
	case "vis":
		fmt.Println(sol.CursorOn)
	case "sample":
		fmt.Printf("%v%v=%v ", sol.Base03, "b03", "Base03")
		fmt.Printf("%v%v=%v ", sol.Base02, "b02", "Base02")
		fmt.Printf("%v%v=%v ", sol.Base01, "b01", "Base01")
		fmt.Printf("%v%v=%v ", sol.Base00, "b00", "Base00")
		fmt.Printf("%v%v=%v ", sol.Base0, "b0", "Base0")
		fmt.Printf("%v%v=%v ", sol.Base1, "b1", "Base1")
		fmt.Printf("%v%v=%v ", sol.Base2, "b2", "Base2")
		fmt.Printf("%v%v=%v\n", sol.Base3, "b3", "Base3")
		fmt.Printf("%v%v=%v ", sol.Yellow, "y", "Yellow")
		fmt.Printf("%v%v=%v ", sol.Orange, "o", "Orange")
		fmt.Printf("%v%v=%v ", sol.Red, "r", "Red")
		fmt.Printf("%v%v=%v ", sol.Magenta, "m", "Magenta")
		fmt.Printf("%v%v=%v ", sol.Violet, "v", "Violet")
		fmt.Printf("%v%v=%v ", sol.Blue, "b", "Blue")
		fmt.Printf("%v%v=%v ", sol.Cyan, "c", "Cyan")
		fmt.Printf("%v%v=%v\n", sol.Green, "g", "Green")
	default:
		usage()
	}
}

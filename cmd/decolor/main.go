package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	color "github.com/ethanbaker/colors"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(color.DecolorAnsiSol(color.DecolorAnsiCss(s.Text())))
	}
	if err := s.Err(); err != nil {
		log.Println(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/skilstak/code/go/color"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(color.Decolor(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Println(err)
	}
}

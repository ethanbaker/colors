package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/ethandbaker/color"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(color.Strip(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Println(err)
	}
}

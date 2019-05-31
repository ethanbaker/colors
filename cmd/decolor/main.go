package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/skilstak/code/colors/go"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(colors.Decolor(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Println(err)
	}
}

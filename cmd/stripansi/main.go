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
		fmt.Println(colors.Strip(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Println(err)
	}
}

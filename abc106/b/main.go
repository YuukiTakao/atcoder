package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	N := readLine()
}

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	r, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return string(r)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	s := scanText()
	// fmt.Printf("%d\n", s)

	rcount := strings.Count(s, "R")
	if rcount == 2 && s[1] == 'S' {
		rcount = 1
	}
	fmt.Printf("%d\n", rcount)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}

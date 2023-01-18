package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Buffer(make([]byte, 1000), 500000)
	sc.Split(bufio.ScanWords)
	_ = scanInt()
	ci := scanText()
	// fmt.Printf("%d %s\n", n, ci)

	rcount := strings.Count(ci, "R")
	wcount := strings.Count(ci, "W")
	if rcount == 0 || wcount == 0 {
		fmt.Printf("0\n")
		return
	}

	ans := 0

	for i := 0; i < rcount; i++ {
		if ci[i] == 'W' {
			ans++
		}
	}

	fmt.Printf("%d\n", ans)
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

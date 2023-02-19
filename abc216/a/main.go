package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	s := scanText()
	// fmt.Printf("%s\n", s)
	last := len(s) - 1
	y := atoi(string(s[last]))
	x := s[0 : last-1]
	if y <= 2 {
		fmt.Printf("%s-\n", x)
	} else if 3 <= y && y <= 6 {
		fmt.Printf("%s\n", x)
	} else if 7 <= y && y <= 9 {
		fmt.Printf("%s+\n", x)
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func atoi(nStr string) int {
	i, err := strconv.Atoi(nStr)
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}

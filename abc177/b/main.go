package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SubstringDistance(s, t string) int {
	distance := len(t)
	for start := 0; start <= len(s)-len(t); start++ {
		diff := 0
		for i := 0; i < len(t); i++ {
			if t[i] != s[i+start] {
				diff++
			}
		}
		if distance > diff {
			distance = diff
		}
	}
	return distance
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	s := scanText()
	t := scanText()
	// fmt.Printf("%s %s\n", s, t)
	fmt.Printf("%d\n", SubstringDistance(s, t))
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

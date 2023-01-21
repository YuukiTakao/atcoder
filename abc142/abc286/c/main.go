package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	a := scanInt()
	b := scanInt()
	s := scanText()
	// fmt.Printf("%d %d %d\n", n, a, b)
	// fmt.Printf("%s\n", s)

	cost := -1
	rcount := 0
	for i := 0; i < n; i++ {
		var sb strings.Builder
		sb.WriteString(s[i:])
		sb.WriteString(string(s[:i]))
		rtStr := sb.String()
		s = rtStr
		rcount++
		fmt.Printf("i=%d rotate=%s s=%s\n", i, rtStr, s)

		var mid int
		if n%2 == 0 {
			mid = n / 2
		} else {
			mid = n/2 + 1
		}
		diff := 0
		for j := 0; j < mid; j++ {
			left := j
			right := n - j - 1

			if s[left] != s[right] {
				diff++
			}
			// fmt.Printf("left=%d right=%d s[left=%c s[right]=%c \n", left, right, s[left], s[right])
		}
		fmt.Printf("%d\n", diff)
		min := minOf(a, b)
		fmt.Printf("%d\n", min*diff)
	}
}

func minOf(vars ...int) int {
	min := 2100000000
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
}

var sc = bufio.NewScanner(os.Stdin)

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

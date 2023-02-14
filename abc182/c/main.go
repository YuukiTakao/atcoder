package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func BitwizeSearch(n int, sets string) int {
	max := -1
	for bit := 0; bit < 1<<n; bit++ {
		pattern := make([]int, 0)
		for i := 0; i < n; i++ {
			if bit&(1<<i) > 0 {
				pattern = append(pattern, i)
			}
		}
		// fmt.Printf("pattern=%v\n", pattern)
		if len(pattern) == 0 {
			continue
		}
		str := bytes.NewBuffer(make([]byte, 0, len(pattern)))
		for _, v := range pattern {
			// fmt.Printf("%c\n", sets[v])
			str.WriteString(string(sets[v]))
		}
		s := str.String()

		num := atoi(s)

		if num%3 == 0 {
			// fmt.Printf("s=%s lens=%d n=%d\n", s, len(s), n)
			if max < len(s) {
				max = len(s)
			}
		}
	}
	if max == -1 {
		return -1
	}
	return n - max
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanText()

	fmt.Printf("%d\n", BitwizeSearch(len(n), n))
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

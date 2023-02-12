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
	t := scanText()

	ans := len(s)
	for start := 0; start <= len(s)-len(t); start++ {

		diff := 0
		for i := 0; i < len(t); i++ {
			if t[i] != s[start+i] {
				diff++
			}
		}
		ans = minOf(ans, diff)
	}

	fmt.Printf("%d\n", ans)
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

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
	h := scanInt()
	w := scanInt()
	// fmt.Printf("%d %d\n", h, w)

	ss := make([]string, h)
	for i := 0; i < h; i++ {
		ss[i] = scanText()
	}

	for i := 0; i < w; i++ {
		sum := 0
		for j := 0; j < h; j++ {
			if ss[j][i] == '#' {
				sum++
			}
		}
		fmt.Printf("%d ", sum)
	}
	fmt.Printf("\n")
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func BitwizeSearch(n, m int, sets []uint) int {
	ans := 0
	for bit := 0; bit < 1<<n; bit++ {
		pattern := make([]int, 0)
		for i := 0; i < n; i++ {
			if bit&(1<<i) > 0 {
				pattern = append(pattern, i)
			}
		}
		// fmt.Printf("pattern=%v\n", pattern)
		bits := uint(0)
		for _, v := range pattern {
			bits |= sets[v]
		}
		if bits == (1<<m)-1 {
			ans++
		}
	}
	return ans
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()

	a := make([]uint, m)
	for i := 0; i < m; i++ {
		c := scanInt()

		for j := 0; j < c; j++ {
			a[i] |= 1 << (scanInt() - 1)
		}
	}

	fmt.Printf("%d\n", BitwizeSearch(m, n, a))
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

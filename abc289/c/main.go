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
	n := scanInt()
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	a := make([]uint, m)
	for i := 0; i < m; i++ {
		c := scanInt()
		bit := uint(0)
		for j := 0; j < c; j++ {
			bit |= 1 << (scanInt() - 1)
		}
		a[i] = bit
		// fmt.Printf("a[i]=%b\n", a[i])
	}

	// fmt.Printf(" (1<<n)-1=%b\n", (1<<n)-1)
	// fmt.Printf("1<<n=%d\n", 1<<n)

	ans := 0
	for bit := 0; bit < 1<<m; bit++ {

		ss := make([]int, 0)
		for i := 0; i < m; i++ {
			if bit&(1<<i) > 0 {
				ss = append(ss, i)
			}
		}
		// fmt.Printf("ss=%v\n", ss)

		bits := uint(0)
		for _, v := range ss {
			bits |= a[v]
		}
		if bits == (1<<n)-1 {
			ans++
		}
		// fmt.Printf("bits=%b bits&((1<<n)-1)=%b\n", bits, bits&((1<<n)-1))
	}
	fmt.Printf("%d\n", ans)
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

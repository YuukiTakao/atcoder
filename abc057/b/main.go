package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()
	}
	// fmt.Printf("%d %d\n", a, b)

	c := make([]int, m)
	d := make([]int, m)
	for i := 0; i < m; i++ {
		c[i], d[i] = scanInt(), scanInt()
	}
	// fmt.Printf("%d %d\n", c, d)

	ans := make([]int, n)
	mins := make([]int, n)
	for i := 0; i < n; i++ {
		mins[i] = 10e10
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// fmt.Printf("abs(a[i])-abs(c[i])=%d\n", abs(a[i])-abs(c[i])+abs(b[j])-abs(d[j]))
			// fmt.Printf("abs(b[i])-abs(d[i])=%d\n", )
			tmp := abs(a[i]-c[j]) + abs(b[i]-d[j])
			// fmt.Printf("tmp=%d\n", tmp)
			if mins[i] > tmp {
				mins[i] = tmp
				ans[i] = j
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", ans[i]+1)
	}
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

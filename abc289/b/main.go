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

	a := make([]int, n)
	for i := 0; i < m; i++ {
		a[i] = scanInt()
	}

	g := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ok := false
		for j := 0; j < m; j++ {
			if i == a[j] {
				ok = true
				break
			}
		}
		if ok {
			g[i] = 1
		} else {
			g[i] = 0
		}
	}
	// fmt.Printf("g=%v\n", g)

	paths := make([][]int, 0, n)
	before := 1
	for i := 1; i <= n; i++ {
		// fmt.Printf("g[%d]=%d\n", i, g[i])
		if g[i] == 0 {
			path := make([]int, 0)
			for k := before; k <= i; k++ {
				path = append(path, k)
			}
			paths = append(paths, path)
			before = i + 1
		}
	}

	for _, v := range paths {
		// fmt.Printf("v=%v\n", v)
		for i := len(v) - 1; i >= 0; i-- {
			fmt.Printf("%d ", v[i])
		}
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

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
	// fmt.Printf("%d\n", n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scanInt()
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		b[i] = scanInt()
	}
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		c[i] = scanInt()
	}

	d := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		d[b[c[i]]]++
	}

	ans := 0
	for _, v := range a {
		ans += d[v]
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

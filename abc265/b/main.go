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
	t := scanInt()
	// fmt.Printf("%d %d %d\n", n, m, t)

	a := make([]int, n+1)
	for i := 1; i <= n-1; i++ {
		a[i] = scanInt()
	}
	x, y := make([]int, m+1), make([]int, m+1)
	bonus := make([]int, n+1)
	for i := 1; i <= m; i++ {
		x[i], y[i] = scanInt(), scanInt()
		bonus[x[i]] = y[i]
	}

	for i := 1; i <= n-1; i++ {
		t -= a[i]
		if t <= 0 {
			fmt.Printf("No\n")
			return
		}
		t += bonus[i+1]
	}
	fmt.Printf("Yes\n")
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInts2(n int) ([]int, []int) {
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()
	}
	return a, b
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

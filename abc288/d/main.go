package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	a := scanInts(n)

	csum := make([][]int, k)
	for i := 0; i < k; i++ {
		csum[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		csum[i%k][i+1] = a[i]
	}
	for j := 0; j < k; j++ {
		for i := 0; i < n; i++ {
			csum[j][i+1] += csum[j][i]
		}
	}

	q := scanInt()
	for qi := 0; qi < q; qi++ {
		l, r := scanInt(), scanInt()
		l--
		// fmt.Printf("l=%d r=%d\n", l, r)
		ns := make([]int, k)
		for i := 0; i < k; i++ {
			ns[i] = csum[i][r] - csum[i][l]
			// fmt.Printf("ns[%d]=%d\n", i, ns[i])
		}
		sort.Ints(ns)
		if ns[0] == ns[k-1] {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}
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

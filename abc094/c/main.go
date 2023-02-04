package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 4096), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	x := scanInts(n)

	sorted := make([]int, n)
	copy(sorted, x)
	sort.Ints(sorted)

	l, r := sorted[(n-1)/2], sorted[n/2]
	for _, v := range x {
		if v <= l {
			fmt.Println(r)
		} else {
			fmt.Println(l)
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

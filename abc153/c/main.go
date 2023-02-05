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

	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = scanInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(h)))
	sum := 0
	for i := 0; i < n; i++ {
		if i < k {
			continue
		}
		sum += h[i]
	}
	fmt.Printf("%d\n", sum)
}
func maxOf(vars ...int) int {
	max := -2100000000
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
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

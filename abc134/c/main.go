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
	// fmt.Printf("%d\n", n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()

	}
	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)

	max := b[n-1]
	subMax := b[n-2]
	for i := 0; i < n; i++ {
		if a[i] == max {
			fmt.Printf("%d\n", subMax)
		} else {
			fmt.Printf("%d\n", max)
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

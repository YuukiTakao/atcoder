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
	a := scanInts(n)
	sort.Ints(a)
	counts := make(map[int]int, n)

	for i := 0; i < n; i++ {
		counts[a[i]]++
	}

	nC2 := n * (n - 1) / 2
	subCnt := 0
	for _, v := range counts {
		if v >= 2 {
			subCnt += v * (v - 1) / 2
		}
	}
	// fmt.Printf("%v %d\n", counts, nC2-subCnt)

	fmt.Printf("%d\n", nC2-subCnt)
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

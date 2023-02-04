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
	x := scanInt()
	// fmt.Printf("%d %d %d\n", n, m, x)

	a := make(map[int]bool) // 料金所のコスト一覧
	for i := 0; i < m; i++ {
		tmp := scanInt()
		a[tmp] = true
	}

	scost := 0
	for i := x - 1; i >= 1; i-- {
		if _, ok := a[i]; ok {
			scost++
		}
	}

	gcost := 0
	for i := x + 1; i < n; i++ {
		if _, ok := a[i]; ok {
			gcost++
		}
	}

	fmt.Printf("%d\n", minOf(scost, gcost))
}

func minOf(vars ...int) int {
	min := 2100000000
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
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

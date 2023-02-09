package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func combinations(n, m int, isUnique bool) [][]int {
	var res [][]int
	path := make([]int, n)
	var dfs func(int)
	dfs = func(u int) {
		if u == n {
			tmp := make([]int, n)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		s := path[u-1]
		if isUnique {
			s += 1
		}
		for i := s; i <= m; i++ {
			path[u] = i
			dfs(u + 1)
		}
	}
	for i := 1; i <= m; i++ {
		path[0] = i
		dfs(1)
	}
	return res
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	for _, c := range combinations(n, m, true) {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", c[i])
		}
		fmt.Println()
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

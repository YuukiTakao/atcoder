package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}
func main() {
	bufInit()
	n := scanInt()

	a := make(map[int]int)
	for i := 1; i <= n; i++ {
		p := scanInt()
		if a[i] == 0 {
			a[p]++
		}
	}
	ans := make([]int, 0)
	for i := 1; i <= n; i++ {
		if a[i] == 0 {
			ans = append(ans, i)
		}
	}
	sort.Ints(ans)
	fmt.Printf("%d\n", len(ans))
	// fmt.Printf("ans=%v\n", ans)

	for i := 0; i < len(ans); i++ {
		fmt.Printf("%d ", ans[i])
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

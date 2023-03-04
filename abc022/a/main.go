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
	s := scanInt()
	t := scanInt()
	// fmt.Printf("%d %d %d\n", n, s, t)

	ans := 0
	w := scanInt()
	if s <= w && w <= t {
		ans++
	}
	a := make([]int, n)
	for i := 0; i < n-1; i++ {
		a[i] = scanInt()
		w += a[i]
		if s <= w && w <= t {
			ans++
		}
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

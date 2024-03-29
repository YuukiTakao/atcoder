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
	h := scanInt()
	w := scanInt()
	r := scanInt()
	c := scanInt()

	ans := 0
	if r-1 >= 1 {
		ans++
	}
	if r+1 <= h {
		ans++
	}
	if c-1 >= 1 {
		ans++
	}
	if c+1 <= w {
		ans++
	}

	// fmt.Printf("%d %d %d %d\n", h, w, r, c)
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

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

	p := make([]int, 2*n+2)
	for i := 1; i <= n; i++ {
		a := scanInt()
		p[2*i] = a
		p[2*i+1] = a
	}

	ans := make([]int, 2*n+2)
	for i := 2; i <= 2*n+1; i++ {
		ans[i] = ans[p[i]] + 1
	}
	for i := 1; i <= 2*n+1; i++ {
		fmt.Printf("%d\n", ans[i])
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

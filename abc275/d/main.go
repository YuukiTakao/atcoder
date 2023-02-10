package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var memo map[int]int

func f(k int) int {
	if k == 0 {
		return 1
	}
	if v, ok := memo[k]; ok {
		return v
	}
	memo[k] = f(k/2) + f(k/3)
	return memo[k]
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()

	memo = make(map[int]int)
	fmt.Printf("%d\n", f(n))
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

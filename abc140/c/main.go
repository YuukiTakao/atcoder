package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	b := make([]int, n)
	for i := 0; i < n-1; i++ {
		b[i] = scanInt()
	}
	b[n-1] = int(math.Pow10(18))
	// fmt.Printf("%v\n", b)

	a := make([]int, n)
	a[0] = b[0]
	for i := 0; i+1 < n; i++ {
		a[i+1] = minOf(b[i], b[i+1])
	}
	// fmt.Printf("%v\n", a)

	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Printf("%d\n", sum)
}
func minOf(vars ...int) int {
	min := int(math.Pow10(18))
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

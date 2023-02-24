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
	k := scanInt()
	n := scanInt()
	// fmt.Printf("%d %d\n", k, n)

	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	max := k - a[n-1] + a[0]
	for i := 0; i+1 < n; i++ {
		tmp := a[i+1] - a[i]
		if max < tmp {
			max = tmp
		}
	}
	fmt.Printf("%d\n", k-max)
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

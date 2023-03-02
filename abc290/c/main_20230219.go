package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func minOf(vars ...int) int {
	min := int(math.Pow10(18))
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
}
func Pow(x, y int) int {
	res := int(math.Pow(float64(x), float64(y)))
	return res
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	// fmt.Printf("%d %d\n", n, k)

	// lim := 3 * Pow(10, 5)
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[scanInt()] = true
	}
	for i := 0; i < k; i++ {
		if !m[i] {
			fmt.Printf("%d\n", i)
			return
		}
	}
	fmt.Printf("%d\n", k)
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

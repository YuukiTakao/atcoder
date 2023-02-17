package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func minOf(vars ...int) int {
	min := 2100000000
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	x := scanInt()
	y := scanInt()
	n := scanInt()
	// fmt.Printf("%d %d %d\n", x, y, n)

	if x*3 < y {
		fmt.Printf("%d\n", n*x)
	} else {
		div := n / 3
		rem := n % 3
		ans := div*y + rem*x
		fmt.Printf("%d\n", ans)
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

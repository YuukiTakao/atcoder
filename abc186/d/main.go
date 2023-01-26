package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	sort.Ints(a)
	sum := 0
	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i] * i
		ans -= sum
		sum += a[i]
	}
	fmt.Printf("%d\n", ans)
}
func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

var sc = bufio.NewScanner(os.Stdin)

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

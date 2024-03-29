package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Pair struct {
	first  int
	weight int
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	w := scanInt()

	c := make([]Pair, n)
	for i := 0; i < n; i++ {
		c[i] = Pair{first: scanInt(), weight: scanInt()}
	}
	sort.Slice(c, func(i, j int) bool { return c[i].first < c[j].first })

	ans := 0
	for i := n - 1; i >= 0; i-- {
		tmpw := minOf(c[i].weight, w)
		ans += c[i].first * tmpw
		w -= tmpw
		if w == 0 {
			break
		}
	}
	fmt.Printf("%d\n", ans)
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

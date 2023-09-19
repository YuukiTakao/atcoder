package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		b[i] = scanInt()
	}

	sort.Ints(a)
	sort.Ints(b)
	// fprintf("a=%v\nb=%v\n", a, b)
	ans := int(math.Pow10(18))
	for i := 0; i < n; i++ {
		lo, hi := -1, len(b)-1
		for (hi - lo) > 1 {
			mid := lo + (hi-lo)/2
			if b[mid] >= a[i] {
				hi = mid
			} else {
				lo = mid
			}
			// fprintf("mid=%d lo=%d hi=%d\n", mid, lo, hi)
		}
		ans = minOf(ans, abs(b[hi]-a[i]))
		if 0 <= hi-1 {
			ans = minOf(ans, abs(b[hi-1]-a[i]))
		}
		// fprintf("ans=%d, b[%d]=%d a[%d]=%d\n", ans, hi, b[hi], i, a[i])
	}
	fprintf("%d\n", ans)
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
func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(a ...interface{}) {
	fmt.Fprintln(wr, a...)
}
func fprint(f string, a ...interface{}) {
	fmt.Fprint(wr, a...)
}
func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}
func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInts2(n int) ([]int, []int) {
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()
	}
	return a, b
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func scanFloat() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
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
func scanBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}

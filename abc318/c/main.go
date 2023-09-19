package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(n, d, p int) {
	f := make([]int, n)
	sumZero := 0
	for i := 0; i < n; i++ {
		f[i] = scanInt()
		sumZero += f[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(f)))

	dp := make([]int, n+1)
	dp[0] = sumZero
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] - f[i-1]
	}
	// fprintf("%v\n", dp)
	ans := dp[0]
	j := 1
	i := d
	for {
		if i > n {
			i = n
		}
		sum := dp[i] + (p * j)
		// fprintf("i=%d sum=%d dp[i]=%d ans=%d\n", i, sum, dp[i], ans)
		if sum >= ans {
			// fprintf("break sum%d ans=%d\n", sum, ans)
			break
		}
		ans = sum
		j++
		i += d
	}
	fprintf("%d\n", ans)
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	d := scanInt()
	p := scanInt()
	solve(n, d, p)
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

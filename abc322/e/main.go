package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	k := scanInt()
	p := scanInt()
	c := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = scanInt()
		a[i] = make([]int, k)
		for j := 0; j < k; j++ {
			a[i][j] = scanInt()
		}
	}
	used := make(map[int]bool, n)
	sum := make([]int, k)
	ans := int(math.Pow10(18))
	var dfsPath func(i int)
	dfsPath = func(i int) {
		if used[i] {
			return
		}
		used[i] = true
		isGteP := true
		for j := 0; j < k; j++ {
			sum[j] += a[i][j]
			if sum[j] < p {
				isGteP = false
			}
		}
		if isGteP {
			cost := 0
			for i, v := range used {
				if v {
					cost += c[i]
				}
			}
			ans = minOf(ans, cost)
		} else {
			for j := 0; j < n; j++ {
				dfsPath(j)
			}
		}
		used[i] = false
		for j := 0; j < k; j++ {
			sum[j] -= a[i][j]
		}
	}
	dfsPath(0)
	if ans == int(math.Pow10(18)) {
		fprintln(-1)
	} else {
		fprintf("%v\n", ans)
	}

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

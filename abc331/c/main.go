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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	counts := make(map[int]int)
	for _, v := range a {
		counts[v]++
	}
	ans := make(map[int]int)
	reversed := deepCopySlice(a)
	sort.Sort(sort.Reverse(sort.IntSlice(reversed)))
	now := 0
	for _, v := range reversed {
		cnt := counts[v]
		if _, ok := ans[v]; !ok {
			ans[v] = now
			now += v * cnt
		}
	}
	for _, v := range a {
		fprintf("%d ", ans[v])
	}
}

func deepCopySlice(s []int) []int {
	tmp := make([]int, len(s))
	copy(tmp, s)
	return tmp
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
	const startBufSize = 4096
	const maxScanTokenSize = math.MaxInt64
	sc.Buffer(make([]byte, startBufSize), maxScanTokenSize)
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

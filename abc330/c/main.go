package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Pow(x, y int) int {
	res := int(math.Pow(float64(x), float64(y)))
	return res
}
func main() {
	bufInit()
	defer wr.Flush()
	d := scanInt()
	searchUpper := int(math.Sqrt(float64(d))) + 1
	ans := searchUpper
	for x := 0; x <= searchUpper; x++ {
		xs := x * x
		y := int(math.Sqrt(float64(d - (x * x))))
		yPlusOne := y + 1
		// fprintf("y=%d xs=%d ans=%d ", y, xs, ans)
		ans = minOf(abs(xs+(y*y)-d), ans)
		ans = minOf(abs(xs+(yPlusOne*yPlusOne)-d), ans)
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

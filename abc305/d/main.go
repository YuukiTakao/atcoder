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

	a := make(map[int]bool, n)
	max := 0
	for i := 0; i < n; i++ {
		in := scanInt()
		a[in] = true
		if max < in {
			max = in
		}
	}
	csum := make([]int, max+1)
	isSleep := false
	csum[0] = 0
	for i := 1; i <= max; i++ {
		if a[i] {
			isSleep = !isSleep
		}
		if max != i && isSleep {
			csum[i] += csum[i-1] + 1
		} else {
			csum[i] += csum[i-1]
		}
		// fprintf("csum[%d]=%d, ", i, csum[i])
	}
	// fprintf("%v\n", csum)
	q := scanInt()

	for i := 0; i < q; i++ {
		l, r := scanInt(), scanInt()
		// fprintf("l=%d r=%d\n", l, r)
		// l = maxOf(l-1, 0)
		// fprintf("%d %d\n", csum[r], csum[l])
		fprintln(csum[r] - csum[l])
	}
}

func maxOf(vars ...int) int {
	max := int(math.Pow10(18)) * -1
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
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
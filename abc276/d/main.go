package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func GCD(a, b int) int {
	if a == 0 || b == 0 {
		if a < b {
			return b
		} else {
			return a
		}
	}
	if a < b {
		a, b = b, a
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}
func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	g := 0
	for i := 0; i < n; i++ {
		g = GCD(g, a[i])
		// fprintf("g=%d\n", g)
	}
	ans := 0
	for i := 0; i < n; i++ {
		// fprintf("a[%d]=%d\n", i, a[i])
		a[i] /= g
		for a[i]%2 == 0 {
			a[i] /= 2
			ans++
		}
		for a[i]%3 == 0 {
			a[i] /= 3
			ans++
		}
		if a[i] != 1 {
			fprintln(-1)
			return
		}
	}
	fprintln(ans)
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

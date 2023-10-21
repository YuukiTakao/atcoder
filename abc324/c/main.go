package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ham(s, t string) int {
	if len(s) != len(t) {
		return 999
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			res++
		}
	}
	return res
}

func f(s, t string) bool {
	if len(s) != len(t)+1 {
		return false
	}
	si := 0
	for ti := 0; ti < len(t); ti++ {
		for si < len(s) && s[si] != t[ti] {
			si++
		}
		if si == len(s) {
			return false
		}
		si++
	}
	return true
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	t := scanText()
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		s := scanText()
		ok := false
		if len(s)+1 >= len(t) {
			if f(s, t) {
				ok = true
			}
			if f(t, s) {
				ok = true
			}
			if ham(s, t) <= 1 {
				ok = true
			}
		}
		if ok {
			ans = append(ans, i+1)
		}
	}

	fmt.Println(len(ans))
	for _, i := range ans {
		fmt.Printf("%d ", i)
	}
	if len(ans) > 0 {
		fmt.Println()
	}
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

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
	m := scanInt()
	s := scanText()
	t := scanText()

	sm := make(map[byte]bool)
	for _, v := range s {
		sm[byte(v)] = true
	}
	// fprintf("%v %s %s\n", sm, s, t)
	for _, v := range t {
		// fprintln(i, v)
		if !sm[byte(v)] {
			fprintln("No")
			return
		}
	}

	if t[0] != s[0] || t[len(t)-1] != s[len(s)-1] {
		fprintln("No")
		return
	}

	substr := make(map[string]bool)
	for i := 0; i+1 < m; i++ {
		sub := t[i : i+2]
		substr[sub] = true
	}
	// fprintln("==", substr)
	for i := 0; i+1 < n; i++ {
		if s[i] != s[0] && s[i+1] != s[n-1] && !substr[s[i:i+2]] {
			fprintln("No")
			return
		}
	}

	exists := false
	for i := 0; i < n; i++ {

		if m+i <= n && t == s[i:m+i] {
			exists = true
		}
		// fprintln("=", t, s[i:m+i], exists)
	}
	if exists {
		fprintln("Yes")
	} else {
		fprintln("No")
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

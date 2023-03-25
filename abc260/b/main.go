package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func (p Pairs) Len() int      { return len(p) }
func (p Pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Pairs) Less(i, j int) bool {
	if p[i].first == p[j].first {
		return p[i].second < p[j].second
	}
	return p[i].first > p[j].first
}

type Pairs []Pair
type Pair struct {
	first  int
	second int
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	x := scanInt()
	y := scanInt()
	z := scanInt()

	a := make(Pairs, n)
	for i := 0; i < n; i++ {
		a[i].first = scanInt()
		a[i].second = i
	}
	b := make(Pairs, n)
	for i := 0; i < n; i++ {
		b[i].first = scanInt()
		b[i].second = i
	}

	passed := make([]bool, n)

	var f func(sl Pairs, lim int)
	f = func(sl Pairs, lim int) {
		tmp := make(Pairs, len(sl))
		copy(tmp, sl)
		sort.Slice(tmp, tmp.Less)
		cnt := 0
		for _, v := range tmp {
			if cnt < lim && !passed[v.second] {
				passed[v.second] = true
				cnt++
			}
		}
	}
	f(a, x)
	f(b, y)
	for i := 0; i < n; i++ {
		a[i].first += b[i].first
	}
	f(a, z)
	for i := 0; i < n; i++ {
		if passed[i] {
			fprintln(i + 1)
		}
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

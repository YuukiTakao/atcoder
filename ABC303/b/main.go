package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type key struct {
	k1, k2 int
}
type tdmap map[key]bool

func (t tdmap) insert(k1, k2 int, v bool) {
	t[key{k1, k2}] = v
}
func (t tdmap) at(k1, k2 int) bool {
	return t[key{k1, k2}]
}
func (t tdmap) erase(k1, k2 int) {
	delete(t, key{k1, k2})
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()

	pictures := make(tdmap, n*m)
	for j := 0; j < m; j++ {
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = scanInt()
			if i >= 1 {
				pictures.insert(a[i-1], a[i], true)
				pictures.insert(a[i], a[i-1], true)
			}
		}
	}

	fprintf("%v\n", pictures)

	ans := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if !pictures.at(i, j) {
				ans++
				// fprintf("%d %d\n", i, j)
			}
			fprintf("i=%d j=%d\n", i, j)
		}
	}
	fprintln(ans)
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Poems []Poem

func (a Poems) Len() int           { return len(a) }
func (a Poems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Poems) Less(i, j int) bool { return a[i].point > a[j].point }

type Poem struct {
	name  string
	point int
	count int
}

func main() {
	bufInit()
	defer wr.Flush()

	n := scanInt()

	counts := make(map[string]int, n)
	p := make(Poems, n)
	for i := 0; i < n; i++ {
		p[i].name = scanText()
		p[i].point = scanInt()
		counts[p[i].name]++
		p[i].count = counts[p[i].name]
	}

	ans := 0
	max := 0
	for i := 0; i < n; i++ {
		if p[i].count == 1 && p[i].point > max {
			max = p[i].point
			ans = i + 1
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
